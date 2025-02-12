package main

import (
	"fmt"
	"github/top1ordie/backen_chat/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type handleAuth func(w http.ResponseWriter, r *http.Request,userP models.User)

func (db *DbCfg) MiddleWareAuth(handle handleAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Cookie")
    tokenString = strings.Replace(tokenString,"Authorization=","",1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			log.Println("token parse error", err)
		}
		var userNickname string
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				log.Println("Expaired token")
				RespondWithJson(w, 403, Error{"Expaired token pls relogin"})
				return
			}
			userNickname = claims["nickname"].(string)
		}

		user, err := db.DB.GetUserByNickName(r.Context(), userNickname)
    retUser := models.DbUserToUser(user)
    retUser.DeletPassword()
    if user.ID == 0 {
      RespondWithJson(w,400,Error{"token not valid " + userNickname})
      return
    }

		//RespondWithJson(w, 200, retUser)
		log.Printf("INFO: token : %v", token)
		handle(w, r,retUser)
	}
}
