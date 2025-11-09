package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BhavneetSingh7/chat/user"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var VALID_ORIGINS = []string{
	"http://127.0.0.1:8080", "http://localhost:8080", "http://127.0.0.1:5500", 
	"http://localhost:5500",
}

var STATIC_FILES = "/static"


func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("GET /static/{loc...}", Static)
	http.HandleFunc("/chat", Chat)
	http.HandleFunc("POST /signup", user.SignUp)
	http.HandleFunc("POST /login", user.Login)
	http.HandleFunc("PATCH /user", user.UpdateEmail)
	http.HandleFunc("DELETE /user", user.RemoveUser)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
