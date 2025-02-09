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

type dbCfg struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Cannot read .env file", err)
	}
	DB_URL := os.Getenv("DB_URL")
	//databaseApi := DatabaseLoad(DB_URL) 
  log.Println(DB_URL)

	r := chi.NewRouter()
  r.Get("/",func(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("Welcome")) 
  })
  http.ListenAndServe(":8080",r)
}

func DatabaseLoad(url string) dbCfg {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln("DB conn error", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatalln("DB ping error", err)
	}
	log.Println("DB connected successfully")
	return dbCfg{DB: database.New(conn)}
}
