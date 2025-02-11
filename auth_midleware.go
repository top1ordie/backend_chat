package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type handleAuth func(w http.ResponseWriter, r *http.Request)

func (db *DbCfg)MiddleWareAuth(handle handleAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Cookie")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        RespondWithJson(w,400,Error{tokenString})
				return nil, fmt.Errorf("Unexpected signing method")

			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			log.Fatalln("token parse error", err)
		}
    var userNickname string
		if claims, ok := token.Claims.(jwt.MapClaims); !ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				log.Fatalln("Expaired token")
        RespondWithJson(w,403,Error{"Expaired token"})
        return
			}
      userNickname = claims["nickname"].(string)
      

		}

    user,err := db.DB.GetUserByNickName(r.Context(),userNickname)

		RespondWithJson(w, 200, user)
		log.Printf("INFO: token : %v", token)
		handle(w, r)
	}
}
