package main

import (
	"encoding/json"
	"github/top1ordie/backen_chat/internal/database"
	"log"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (db *DbCfg) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "BAD DATA", 400)
		return
	}

	hashpassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatalln("Hash password error :", err)
		return
	}
	tokenString, err := GenerateJWT(user.Nickname)
  if err!= nil {
    log.Fatalln("JWT ",err)
  }
	http.SetCookie(w, &http.Cookie{
		Name:        "Authorization",
		Value:       tokenString,
		Quoted:      false,
		Path:        "",
		Domain:      "",
		Expires:     time.Time{},
		RawExpires:  "",
		MaxAge:      0,
		Secure:      false,
		HttpOnly:    false,
		SameSite:    0,
		Partitioned: false,
		Raw:         "",
		Unparsed:    []string{},
	})

	_, err = db.DB.CreateUserNoId(r.Context(), database.CreateUserNoIdParams{Nickname: user.Nickname, Password: hashpassword})
	if err != nil {
		log.Fatalln("Cannot create new user in DB", err)
		return
	}
	http.StatusText(200)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
