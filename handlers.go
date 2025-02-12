package main

import (
	"encoding/json"
	"github/top1ordie/backen_chat/internal/database"
	"github/top1ordie/backen_chat/models"
	"log"
	"net/http"
)

func (db *DbCfg) AddUserToChat(w http.ResponseWriter, r *http.Request, user models.User) {
	decoder := json.NewDecoder(r.Body)
	var userAddGroup models.UserAddGroup
	if err := decoder.Decode(&userAddGroup); err != nil {
		RespondWithJson(w, 400, Error{"malformed json"})
		log.Println("malformed json: ", err)
		return
	}

	log.Printf("chat id %d user id %d", userAddGroup.IdChat, userAddGroup.IdUser)

	_, err := db.DB.CreateUser_Chat(r.Context(), database.CreateUser_ChatParams{
		UserID: int32(userAddGroup.IdUser),
		ChatID: int32(userAddGroup.IdChat),
	})
	if err != nil {
		RespondWithJson(w, 500, Error{"Database error"})
		log.Println("ERROR : database error cannot add user ", err)
		return
	}

	RespondWithJson(w, 200, Info{"Succes"})
}
