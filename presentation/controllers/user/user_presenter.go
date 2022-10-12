package controllers

import (
	"clean_arch_example/domain/entities"
	"clean_arch_example/presentation"
)

func MakeUserPresenter() presentation.Presenter[entities.User] {
	return &userPresenter{}
}

type userPresenter struct{}

func (presenter *userPresenter) Parse(user entities.User) map[string]any {
	return map[string]any{"id": user.ID, "name": user.Name}
}
