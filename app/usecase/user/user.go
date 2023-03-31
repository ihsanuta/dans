package user

import (
	"dans/app/model"
	"dans/app/repository/user"
	"dans/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserUsecase interface {
	Login(payload model.Login) (model.LoginResponse, error)
}

type userUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(userRepository user.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Login(payload model.Login) (model.LoginResponse, error) {
	resp := model.LoginResponse{}
	user, err := u.userRepository.Login(payload)
	if err != nil {
		return resp, err
	}
	resp.Token, err = u.GenerateTokenJWT(user)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *userUsecase) GenerateTokenJWT(user model.User) (string, error) {
	secretKey := config.JwtConfig["jwt_signature"].(string)
	if secretKey == "" {
		return "", fmt.Errorf("%s", "error get signature key")
	}

	exp := time.Now().AddDate(1, 0, 0)
	issue := time.Now()
	claims := model.TokenData{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(issue),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// encoded string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("%s", "error generate token")
	}

	return tokenString, nil
}
