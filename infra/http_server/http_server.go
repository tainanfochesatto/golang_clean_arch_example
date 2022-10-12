package httpServer

type HttpServer interface {
	On(route string, method string, handler func(body HttpRequest) HttpResponse)
	Listen(port string)
}
