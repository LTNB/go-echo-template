package i18n

import (
	"fmt"
	hocon "github.com/go-akka/configuration"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const defaultLocale = "vi-vn"

var I18 I18n

/*
 * init i18n witt ${defaultLocale}
 */
func NewI18n(dir string) {
	log.Printf("Loading i18n files from directory [%s]", dir)

	textFromAllFiles := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) != ".i18n" {
			return nil
		}
		log.Printf("\tLoading i18n file [%s]...", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		textFromAllFiles = append(textFromAllFiles, string(data))
		return nil
	})
	if err != nil {
		panic(err)
	}
	i18n := I18n{config: hocon.ParseString(strings.Join(textFromAllFiles, "\n"))}
	i18n.locale = defaultLocale
	I18 = i18n
}

// TODO: change locale?
type I18n struct {
	config *hocon.Config
	locale string
}

func (i18n *I18n) SetLocale(locale interface{}) {
	if locale != nil && locale.(string) != "" {
		i18n.locale = locale.(string)
	} else {
		i18n.locale = defaultLocale
	}

}

func (i18n I18n) FlashMsg(locale, path string, params ...interface{}) string {
	if locale == "" {
		locale = i18n.locale
	}
	path = locale + ".text." + path
	format := i18n.config.GetString(path, path)
	return fmt.Sprintf(format, params...)
}

func (i18n I18n) Text(path string, params ...interface{}) string {
	if i18n.locale != "" {
		path = i18n.locale + ".text." + path
	}
	format := i18n.config.GetString(path, path)
	return fmt.Sprintf(format, params...)
}
