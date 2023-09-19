package user

import "go-restaurant/m/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegisetered(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
}
