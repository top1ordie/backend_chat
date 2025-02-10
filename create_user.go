package main

import (
	"encoding/json"
	"github/top1ordie/backen_chat/internal/database"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
  }
	_, err = db.DB.CreateUserNoId(r.Context(), database.CreateUserNoIdParams{Nickname: user.Nickname, Password: hashpassword})
	if err != nil {
		log.Fatalln("Cannot create new user in DB", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
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
	http.StatusText(200)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
