package echo_init

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	config "main/src/init"
	"main/src/init/i18n"
	"strings"
)

// Ref: https://echo.labstack.com/guide/templates

/*
 * templateRenderer is a custom html/template renderer for Echo framework
 */
type templateRenderer struct {
	directory          string
	templateFileSuffix string
	templates          map[string]*template.Template
	i18n               i18n.I18n
}

func newTemplateRenderer(directory, templateFileSuffix string, i18n i18n.I18n) *templateRenderer {
	return &templateRenderer{
		directory:          directory,
		templateFileSuffix: templateFileSuffix,
		templates:          map[string]*template.Template{},
		i18n:               i18n,
	}
}

// Render renders a template document
func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	sess := GetSession(c)
	flash := sess.Flashes()
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	// add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		t.i18n.SetLocale(sess.Values["locale"])

		viewContext["i18n"] = t.i18n
		viewContext["reverse"] = c.Echo().Reverse
		viewContext["static"] = staticPath
		viewContext["appInfo"] = config.AppConfig.Conf.GetConfig("app")

		if len(flash) > 0 {
			viewContext["flash"] = flash[0].(string)
		}
	}
	tpl := t.templates[name]
	tokens := strings.Split(name, ":")
	if tpl == nil {
		var files []string
		for _, v := range tokens {
			files = append(files, t.directory+"/"+v+".html")
		}
		tpl = template.Must(template.New(name).ParseFiles(files...))
		t.templates[name] = tpl
	}

	return tpl.ExecuteTemplate(w, tokens[0]+".html", data)
}
