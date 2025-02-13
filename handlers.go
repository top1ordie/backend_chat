package main

import (
	"encoding/json"
	"github/top1ordie/backen_chat/internal/database"
	"github/top1ordie/backen_chat/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (db *DbCfg) AddUserToChat(w http.ResponseWriter, r *http.Request, user models.User) {
	decoder := json.NewDecoder(r.Body)
	var userAddGroup models.User_Chat
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

func (db *DbCfg) CreateMessage(w http.ResponseWriter, r *http.Request, user models.User) {
	decoder := json.NewDecoder(r.Body)
	var Messages models.Message
	err := decoder.Decode(&Messages)
	if err != nil {
		RespondWithJson(w, 400, Error{"malformed json Mesagges"})
		return
	}
	userR, err := db.DB.GetUserInChat(r.Context(), database.GetUserInChatParams{
		ID:   int32(Messages.User_Chat.IdUser),
		ID_2: int32(user.Id),
	})
	if err != nil {
		RespondWithJson(w, 500, Error{"DB error"})
    log.Printf("ERROR : %v \n",err)
    log.Printf("ERROR : useId: %d, chatId: %d\n",user.Id,Messages.User_Chat.IdChat)
		return
	}
	if userR.ID == 0 {
		RespondWithJson(w, 400, Error{"User doesnt have permision to send message for this chats"})
		return
	}

  dbMessage, err := db.DB.CreateMessage(r.Context(), database.CreateMessageParams{
		ChatID:    int32(Messages.User_Chat.IdChat),
		UserID:    int32(user.Id),
		CreatedAt: time.Now(),
	})
	if err != nil {
		RespondWithJson(w, 500, Error{"database error"})
		return
	}
  for _,i := range Messages.MessageArr {
    _, err = db.DB.CreateTextMessage(r.Context(),database.CreateTextMessageParams{
    	MessageID: dbMessage,
    	Data:      i.Data,
    })
    if err!= nil {
      RespondWithJson(w,500,"DATABASE err ")
      return
    }
  }


	RespondWithJson(w, 200, Info{strconv.Itoa(len(Messages.MessageArr))})
}
