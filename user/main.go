package user

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Email string
	Password string
	CreatedAt time.Time
	ModifiedAt time.Time
}

var Users map[string]map[string]string = map[string]map[string]string{}

func CreateUser(user_data UserData) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(user_data.Password), 0)
	if err != nil {
		log.Println("error occured while generating password hash", err)
		return err
	}

	if Users[user_data.Email] != nil {
		return errors.New("user already exists")
	}

	user_data.CreatedAt = time.Now().UTC()
	user_data.ModifiedAt = user_data.CreatedAt

	Users[user_data.Email] = map[string]string{
		"email": user_data.Email,
		"password": string(pwd),
		"created_at": user_data.CreatedAt.GoString(),
		"modified_at": user_data.ModifiedAt.GoString(),
	}

	return nil
}


func GetUser(email string) (map[string]string, error) {
	lookup := Users[email]
	if lookup == nil {
		return map[string]string{}, errors.New("user not found")
	}
	return lookup, nil
}


func UpdateUser(email, password string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Println("error occured while generating password hash", err)
		return err
	}

	if Users[email] == nil {
		return errors.New("user not found")
	}
	current := Users[email]
	Users[email] = map[string]string{
		"email": email,
		"password": string(pwd),
		"created_at": current["created_at"],
		"modified_at": time.Now().UTC().GoString(),
	}

	return nil
}



func DeleteUser(email string) error {
	if Users[email] == nil {
		return errors.New("user not found")
	}
	Users[email] = map[string]string{}
	return nil
}

