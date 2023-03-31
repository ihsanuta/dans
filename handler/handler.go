package handler

import (
	usecase "dans/app/usecase"
	"dans/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Handler interface{}

var once = &sync.Once{}

type handler struct {
	usecase *usecase.Usecase
}

func Init(usecase *usecase.Usecase) *handler {
	var h *handler
	once.Do(func() {
		h = &handler{
			usecase: usecase,
		}
		h.Serve()
	})
	return h
}

func (h *handler) Serve() {
	router := gin.Default()
	group := router.Group("/api/v1")
	// group.GET("/articles", c.GetArticles)
	group.POST("/login", h.Login)
	group.GET("/job/:job_id", h.authenticateToken, h.GetByID)
	group.GET("/job", h.authenticateToken, h.List)

	serverString := fmt.Sprintf("%s:%s", config.AppConfig["host"], config.AppConfig["port"])
	router.Run(serverString)
}

func (h *handler) authenticateToken(c *gin.Context) {
	tokenString := GetTokenFromGinContext(c)

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		c.Abort()
		return
	}
	err := ValidateToken(tokenString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}

func GetTokenFromGinContext(c *gin.Context) string {
	authorizationHeader := c.GetHeader("Authorization")

	authorizationValues := strings.SplitN(authorizationHeader, " ", 2)

	if len(authorizationValues) < 2 || strings.ToLower(authorizationValues[0]) != "bearer" {
		return ""
	}

	return strings.TrimSpace(authorizationValues[1])
}

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func ValidateToken(signedToken string) (err error) {
	secretKey := config.JwtConfig["jwt_signature"].(string)
	if secretKey == "" {
		return fmt.Errorf("%s", "error get signature key")
	}

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
