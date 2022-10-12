package controllers

import (
	"clean_arch_example/domain/entities"
	"clean_arch_example/domain/usecases"
	"clean_arch_example/presentation"
	"log"
)

type UserController struct {
	CreateUser usecases.CreateUser
	FindUser   usecases.FindUser
	Presenter  presentation.Presenter[entities.User]
}

type UserDto struct{ Name string }

func (controller UserController) Create(dto UserDto) map[string]any {
	user, err := controller.CreateUser.Call(dto.Name)
	if err != nil {
		log.Fatal(err.Error())
	}
	return controller.Presenter.Parse(*user)
}

func (controller UserController) Find(id string) map[string]any {
	user, err := controller.FindUser.Call(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	return controller.Presenter.Parse(*user)
}
