package httpServer

type HttpRequest struct {
	Body   map[string]interface{}
	Params map[string]string
}
