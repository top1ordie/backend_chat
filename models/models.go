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

type UserAddGroup struct {
  IdUser int `json:"id_user"`
  IdChat int `json:"id_chat"`
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
