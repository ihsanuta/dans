package user

import (
	"dans/app/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Login(payload model.Login) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Login(payload model.Login) (model.User, error) {
	var resp model.User

	query := u.db.Raw("SELECT id, username, password FROM users WHERE username = ? AND password = ?", payload.Username, payload.Password).Scan(&resp)

	if query.Error != nil {
		return resp, query.Error
	}

	return resp, nil
}
