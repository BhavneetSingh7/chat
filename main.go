package main

import (
	"flag"
	"log"
	"net/http"
	// "io"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var VALID_ORIGINS = []string{
	"http://127.0.0.1:8080", "http://localhost:8080", "http://127.0.0.1:5500", 
	"http://localhost:5500",
}

var STATIC_FILES = "/static"

func Static(w http.ResponseWriter, r *http.Request) {
	ls := strings.Split(r.URL.Path, STATIC_FILES)
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

	_, err = w.Write(x[:n])
	if err != nil {
		log.Println("error writing message: ", err)
	}
}

func Chat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return slices.Contains(VALID_ORIGINS, origin)
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("error while reading incoming message:", err)
		}
		time.Sleep(500 * time.Millisecond) //Delay
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("error writing message: ", err)
			break
		}
		// log.Println("message type: ", mt)
		// x := make([]byte, 512)
		// f, err := os.Open("data.txt")
		// if err != nil {
		// 	log.Print("error in opening file: ", err)
		// 	return
		// }

		// for {
		// 	_, err := f.Read(x)
		// 	if err != nil {
		// 		if err != io.EOF {log.Print("error in reading file: ", err)}
		// 		break
		// 	}
		// 	time.Sleep(50 * time.Millisecond) //Delay
		// 	err = c.WriteMessage(mt, msg)
		// 	if err != nil {
		// 		log.Println("error writing message: ", err)
		// 		break
		// 	}
			
		// }
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("GET /static/{loc...}", Static)
	http.HandleFunc("/chat", Chat)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
