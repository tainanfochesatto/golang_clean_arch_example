package repositories

import (
	"clean_arch_example/domain/entities"
	"clean_arch_example/domain/repositories"
	"errors"
)

func MakeMemoryUserRepository() repositories.UserRepository {
	return &memoryUserRepository{}
}

type memoryUserRepository struct {
	Users []entities.User
}

func (repository *memoryUserRepository) Create(user entities.User) error {
	repository.Users = append(repository.Users, user)
	return nil
}

func (repository *memoryUserRepository) Find(id string) (*entities.User, error) {
	for i := range repository.Users {
		if repository.Users[i].ID == id {
			return &repository.Users[i], nil
		}
	}
	return nil, errors.New("user not found")
}
