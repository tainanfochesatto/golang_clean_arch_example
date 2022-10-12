package repositories

import (
	"clean_arch_example/domain/entities"
	"clean_arch_example/domain/repositories"
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
