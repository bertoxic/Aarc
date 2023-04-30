package repository

import "github.com/bertoxic/aarc/models"

type DatabaseRepo interface {
	GetUserByEmail(email string) (models.User, error)
    UpdateUser(user models.User) error
    CreateUser(user models.User) error
    IsEmailUsed(email string)(bool, error)
    GetVerifiedUserByEmail(email string) (models.User, error)
    CheckifTableExist(email string) error
}