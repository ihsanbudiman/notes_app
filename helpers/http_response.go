package helpers

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
