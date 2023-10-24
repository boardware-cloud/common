package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

var secret []byte

func Init() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	secret = []byte(viper.GetString("jwt.secret"))
}

type Claims struct {
	jwt.StandardClaims
	ID     string `json:"id"` // Account Id
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func NewClaims(id, email, role string, expiredAt int64, status string) Claims {
	var c Claims
	c.ID = id
	c.Email = email
	c.ExpiresAt = expiredAt
	c.Role = role
	c.Status = status
	return c
}

func SignJwt(id, email, role string, expiredAt int64, status string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(id, email, role, expiredAt, status))
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func VerifyJwt(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("Unvalid token")
	}
}
