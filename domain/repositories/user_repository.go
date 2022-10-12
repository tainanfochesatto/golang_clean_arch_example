package repositories

import entities "clean_arch_example/domain/entities"

type UserRepository interface {
	Create(user entities.User) error
	Find(id string) (*entities.User, error)
}
