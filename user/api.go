package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user_data := UserData{}
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	r.Body.Read(body)

	err := json.Unmarshal(body, &user_data)
	if err != nil {
		log.Println("failed to read request json", err)
		w.WriteHeader(400)
		response, _ := json.Marshal(map[string]string{
			"error": "failed to read request json",
		})
		w.Write(response)
		return
	}

	err = CreateUser(user_data)
	if err != nil {
		log.Println("failed to create user", err)
		w.WriteHeader(400)
		response, _ := json.Marshal(map[string]string{
			"error": fmt.Sprintf("failed to create user: %s", err),
		})
		w.Write(response)
		return
	} else {
		// Return access token and set refresh token in cookies
		w.WriteHeader(202)
		response, _ := json.Marshal(user_data)
		w.Write(response)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	r.Body.Read(body)

	var req_json map[string]string
	err := json.Unmarshal(body, &req_json)
	if err != nil {
		log.Println("failed to read request json", err)
		w.WriteHeader(400)
		response, _ := json.Marshal(map[string]string{
			"error": "failed to read request json",
		})
		w.Write(response)
		return
	}

	user_data, err := GetUser(req_json["email"])
	if err != nil {
		log.Println("user not found: ", err)
		w.WriteHeader(404)
		response, _ := json.Marshal(map[string]string{
			"error": "user not found",
		})
		w.Write(response)
		return
	}

	// Return access token and set refresh token in cookies
	w.WriteHeader(200)
	response, _ := json.Marshal(user_data)
	w.Write(response)
}


func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	r.Body.Read(body)

	var req_json map[string]string
	err := json.Unmarshal(body, &req_json)
	if err != nil {
		log.Println("failed to read request json", err)
		w.WriteHeader(400)
		response, _ := json.Marshal(map[string]string{
			"error": "failed to read request json",
		})
		w.Write(response)
		return
	}

	user_data, err := GetUser(req_json["email"])
	if err != nil {
		log.Println("user not found: ", err)
		w.WriteHeader(404)
		response, _ := json.Marshal(map[string]string{
			"error": "user not found",
		})
		w.Write(response)
		return
	}

	// Return access token and set refresh token in cookies
	w.WriteHeader(200)
	response, _ := json.Marshal(user_data)
	w.Write(response)
}


func RemoveUser(w http.ResponseWriter, r *http.Request) {
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	r.Body.Read(body)

	var req_json map[string]string
	err := json.Unmarshal(body, &req_json)
	if err != nil {
		log.Println("failed to read request json", err)
		w.WriteHeader(400)
		response, _ := json.Marshal(map[string]string{
			"error": "failed to read request json",
		})
		w.Write(response)
		return
	}

	user_data, err := GetUser(req_json["email"])
	if err != nil {
		log.Println("user not found: ", err)
		w.WriteHeader(404)
		response, _ := json.Marshal(map[string]string{
			"error": "user not found",
		})
		w.Write(response)
		return
	}

	// Return access token and set refresh token in cookies
	w.WriteHeader(200)
	response, _ := json.Marshal(user_data)
	w.Write(response)
}

