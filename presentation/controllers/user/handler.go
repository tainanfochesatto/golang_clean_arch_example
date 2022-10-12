package controllers

import (
	"clean_arch_example/data/repositories"
	"clean_arch_example/domain/usecases"
	httpServer "clean_arch_example/infra/http_server"

	"github.com/mitchellh/mapstructure"
)

var userRepository = repositories.MakeMemoryUserRepository()
var createUser = usecases.CreateUser{UserRepository: userRepository}
var userController = UserController{
	CreateUser: createUser,
	Presenter:  MakeUserPresenter(),
}

func CreateUserHandler(request httpServer.HttpRequest) httpServer.HttpResponse {
	var userDto UserDto
	mapstructure.Decode(request.Body, &userDto)
	response := userController.Create(userDto)
	return httpServer.HttpResponse{StatusCode: 201, Body: response}
}
