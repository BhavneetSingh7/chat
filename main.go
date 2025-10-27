package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func static(w http.ResponseWriter, r *http.Request) {
	ls := strings.Split(r.URL.Path, "/static")
	file_path := "." + ls[1]
	f, err := os.Open(file_path)
	if err != nil {
		log.Print("error in opening file: ", err)
		w.WriteHeader(404)
		w.Write([]byte("file not found"))
		return
	}
	defer f.Close()

	x := make([]byte, 8192)
	n, ferr := f.Read(x)
	if ferr != nil {
		log.Print("error in reading file: ", ferr)
		w.WriteHeader(404)
		w.Write([]byte("file not found"))
		return
	}
	// w.Header().Set("Content-Disposition", "inline")
	_, err = w.Write(x[:n])
	if err != nil {
		log.Println("error writing message: ", err)
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	valid_origins := []string{
		"http://127.0.0.1:8080", "http://localhost:8080", "http://127.0.0.1:5500", 
		"http://localhost:5500",
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		// log.Println("Incoming request from origin: ", origin)
		return slices.Contains(valid_origins, origin)
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	mt, _, err := c.ReadMessage()
	if err != nil {
		log.Println("error while reading incoming message:", err)
	}
	// log.Println("message type: ", mt)
	x := make([]byte, 512)
	f, err := os.Open("data.txt")
	if err != nil {
		log.Print("error in opening file: ", err)
		return
	}

	for {
		_, err := f.Read(x)
		if err != nil {
			log.Print("error in reading file: ", err)
			break
		}
		time.Sleep(50 * time.Millisecond) //Delay
		err = c.WriteMessage(mt, x)
		if err != nil {
			log.Println("error writing message: ", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("GET /static/{loc...}", static)
	http.HandleFunc("/chat", data)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
