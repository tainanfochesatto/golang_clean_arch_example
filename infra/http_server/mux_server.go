package httpServer

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeMuxHttpServer() HttpServer {
	return &muxHttpServer{Server: *mux.NewRouter()}
}

type muxHttpServer struct{ Server mux.Router }

func (httpServer *muxHttpServer) On(route string, method string, handler func(body HttpRequest) HttpResponse) {
	httpServer.Server.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		var request HttpRequest
		json.NewDecoder(r.Body).Decode(&request.Body)
		response := handler(request)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response.Body)
	}).Methods(method)
}

func (httpServer *muxHttpServer) Listen(port string) {
	http.ListenAndServe(port, &httpServer.Server)
}
