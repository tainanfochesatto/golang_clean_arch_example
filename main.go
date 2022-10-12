package main

import (
	httpServer "clean_arch_example/infra/http_server"
	controllers "clean_arch_example/presentation/controllers/user"
)

func main() {
	server := httpServer.MakeMuxHttpServer()
	server.On("/users", "POST", controllers.CreateUserHandler)
	server.Listen(":3000")
}
