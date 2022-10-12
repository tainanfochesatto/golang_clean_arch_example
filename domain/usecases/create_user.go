package usecases

import (
	"clean_arch_example/domain/entities"
	"clean_arch_example/domain/repositories"
	"math/rand"
	"strconv"
)

type CreateUser struct {
	UserRepository repositories.UserRepository
}

func (us *CreateUser) Call(name string) (*entities.User, error) {
	user := entities.User{
		ID:   strconv.Itoa(rand.Intn(10000)),
		Name: name,
	}
	err := us.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
