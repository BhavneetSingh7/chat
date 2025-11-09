package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

func Static(w http.ResponseWriter, r *http.Request) {
	ls := strings.Split(r.URL.Path, STATIC_FILES)
	file_path := ls[1]

	x, err := fs.ReadFile(os.DirFS("."), file_path[1:])
	if err != nil {
		log.Print("error in reading file: ", err)
		w.WriteHeader(404)
		w.Write([]byte("file not found"))
		return
	}

	_, err = w.Write(x)
	if err != nil {
		log.Println("error writing message: ", err)
	}
}