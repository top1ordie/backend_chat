package main

import (
	"encoding/json"
	"github/top1ordie/backen_chat/internal/database"
	"github/top1ordie/backen_chat/models"
	"log"
	"net/http"
)

func (db *DbCfg) CreateChat(w http.ResponseWriter, r *http.Request, user models.User) {
	decoder := json.NewDecoder(r.Body)
	var chat models.Chat

	err := decoder.Decode(&chat)
	if err != nil {
		RespondWithJson(w, 400, Error{"malform json"})
		return
	}
	chatDb, err := db.DB.CreateChat(r.Context(), chat.ChatName)
	if err != nil {
		RespondWithJson(w, 400, Error{"cannot create chat"})
		log.Println("cannot create chat", err)
	}
	_, err = db.DB.CreateUser_Chat(r.Context(), database.CreateUser_ChatParams{
		UserID: int32(user.Id),
		ChatID: chatDb.ID,
	})
	if err != nil {
		RespondWithJson(w, 400, Error{"cannote create chat_user relations"})
		log.Printf("cannote create chat_user %v", err)
	}

	RespondWithJson(w, 200, models.DbChatToChat(chatDb))
}
