package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (db *DbCfg) Login_User(w http.ResponseWriter, r *http.Request) {
	var user User
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&user)
	if err != nil {
		log.Println("Decoder error: ", err)
		return
	}
	rUser, err := db.DB.GetUserByNickName(r.Context(), user.Nickname)

	checkPas := CheckPasswordHash(user.Password, rUser.Password)
	if !checkPas {
		log.Println("wrong cred : ", err)
		return
	}

	tokenString, err := GenerateJWT(user.Nickname)
	if err != nil {
		log.Println("JWT", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:        "Authorization",
		Value:       tokenString,
		Quoted:      false,
		Path:        "",
		Domain:      "",
		Expires:     time.Now().Add(time.Hour * 24 * 10),
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
