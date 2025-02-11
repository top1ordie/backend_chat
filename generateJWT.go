package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(nickname string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nickname": nickname,
		"exp":      time.Now().Add(time.Hour * 24 * 10).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatalf("jwt token error : %v", err)
    return "", errors.New("jwt token error : ")
	}
  return tokenString,nil 
}
