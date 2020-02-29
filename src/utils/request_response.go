package utils

type ResponseJson struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}


const (
	statusOk          = 200
	statusForbidden   = 403
	statusNotFound    = 404
	statusConflict    = 409
	statusErrorClient = 400
	statusErrorServer = 500

	CtxKeyResponse = "response"
)
