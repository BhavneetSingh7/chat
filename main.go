package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	x := make([]byte, 512)
	f, err := os.Open("data.txt")
	if err != nil {
		log.Print("error in opening file: ", err)
		return
	}
	mt, _, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}
	for {
		n, err := f.Read(x)
		if err != nil {
			log.Print("error in reading file: ", err)
			break
		}
		if n==0 {
			log.Print("EOF reached.")
			c.Close()
			break
		}
		time.Sleep(50*time.Millisecond)
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
	// http.HandleFunc("/echo", echo)
	http.HandleFunc("/", data)
	log.Fatal(http.ListenAndServe(*addr, nil))
}