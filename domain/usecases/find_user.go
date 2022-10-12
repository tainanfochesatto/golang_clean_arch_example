package usecases

import (
	"clean_arch_example/domain/entities"
	"clean_arch_example/domain/repositories"
)

type FindUser struct {
	UserRepository repositories.UserRepository
}

func (us *FindUser) Call(id string) (*entities.User, error) {
	user, err := us.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
