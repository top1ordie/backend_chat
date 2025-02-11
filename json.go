package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct{
  Error string `json:"error"`
}
type Info struct{
  Info string `json:"info"`
}

func RespondWithJson(w http.ResponseWriter,code int,payload interface{}) {
  data,err:= json.Marshal(payload)
  if err != nil {
    log.Fatalln("payload error")
    w.WriteHeader(500)
  }
  w.Header().Add("Content-Type","application/json")
  w.WriteHeader(code)
  w.Write(data)
}
