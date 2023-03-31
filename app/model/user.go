package model

import "github.com/golang-jwt/jwt/v4"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type TokenData struct {
	Data User `json:"data"`
	jwt.RegisteredClaims
}
