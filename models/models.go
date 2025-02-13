package models

import "github/top1ordie/backen_chat/internal/database"


type User struct {
  Id int `json:"id"`
  Nickname string `json:"nickname"`
  Password string `json:"password"`
}
type Chat struct {
  Id int `json:"id"`
  ChatName string `json:"chat_name"`
}

type User_Chat struct {
  IdUser int `json:"id_user"`
  IdChat int `json:"id_chat"`
}

type Text_Message struct {
  Id int `json:"id"`
  MessageId int `json:"message_id"`
  Data string `json:"data"`
}
type Media_Message struct {
  Id int `json:"id"`
  MessageId int `json:"message_id"`
  Data string `json:"data_url"`
}

type Message struct {
  MessageArr []Text_Message `json:"message_arr"`
  User_Chat User_Chat `json:"chat_user"`
}

func DbUserToUser(dbUser database.User) User{
  return User{
  	Id:       int(dbUser.ID),
  	Nickname: dbUser.Nickname,
  	Password: dbUser.Password,
  }
}
func DbChatToChat(dbChat database.Chat) Chat{
  return Chat{
  	Id:       int(dbChat.ID),
  	ChatName: dbChat.ChatName,
  }
}

func (user *User)DeletPassword() {
  user.Password = ""
}
