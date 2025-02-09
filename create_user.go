package main

import (
	"encoding/json"
	"github/top1ordie/backen_chat/internal/database"
	"log"
	"net/http"
)

type User struct {
	nickname string `json:"nickname"`
	password string `json:"password"`
}

func (db *DbCfg) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
    http.Error(w,"BAD DATA",400)
    return
	}
  _ , err =db.DB.CreateUserNoId(r.Context(),database.CreateUserNoIdParams{Nickname: user.nickname,Password: user.password})
  if err != nil {
    log.Fatalln("Cannot create new user in DB",err)
  }
}
