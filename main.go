package main

import (
	"database/sql"
	"github/top1ordie/backen_chat/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbCfg struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Cannot read .env file", err)
	}
	DB_URL := os.Getenv("DB_URL")
	databaseApi := DatabaseLoad(DB_URL) 
  log.Println(DB_URL)
	r := chi.NewRouter()
  r.Get("/",func(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("Welcome")) 
  })
  r.Post("/signUp",databaseApi.SignUpHandler)
  r.Post("/signIn",databaseApi.Login_User)
  r.Post("/createChat",databaseApi.MiddleWareAuth(databaseApi.CreateChat))
  r.Post("/addUserToChat",databaseApi.MiddleWareAuth(databaseApi.AddUserToChat))
  r.Post("/createMessage",databaseApi.MiddleWareAuth(databaseApi.CreateMessage))
  http.ListenAndServe(":8080",r)
}

func DatabaseLoad(url string)  DbCfg{
	conn, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln("DB conn error", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalln("DB ping error", err)
	}
	log.Println("DB connected successfully")
	return DbCfg{DB: database.New(conn)}
}
