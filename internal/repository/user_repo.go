package repository

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/ArteShow/ArteDataStorage/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CreateUser(username, password string) (error, string) {
	hash, err := hashPassword(password)
	if err != nil{
		return err, ""
	}

	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return err, ""
	}

	id := hex.EncodeToString(bytes)

	database.Connect()
	defer database.Close()

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT FROM users WHERE password = ? AND username = ?)", password, username).Scan(&exists)
	if err != nil{
		return err, ""
	} else if !exists{
		return errors.New("User with the same username or password already exists"), ""
	}

	_, err2 := database.DB.Exec(
		"INSERT INTO users (id, username, password) VALUES (?, ?, ?)",
		id, username, hash,
	)

	if err2 != nil {
		return err2, ""
	}

	return nil, id
}

func DeleteUser(id string) error { 
	database.Connect()
	defer database.Close()

	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil{
		return err
	}

	return nil
}

func UpdateUsername(id, username string) error{
	database.Connect()
	defer database.Close()

	_, err := database.DB.Exec("UPDATE users SET username = ? WHERE id = ?", username, id)
	if err != nil{
		return err
	}

	return nil
}

func UpdateUserPassword(id, password, newPassword string) error{
	database.Connect()
	defer database.Close()

	hashedPassword, err := hashPassword(password)
	if err != nil{
		return err
	}
	
	hashNewPassword, err := hashPassword(newPassword)
	if err != nil{
		return err
	}

	var currentPassword string
	err = database.DB.QueryRow("SELECT password FROM users WHERE id = ?", id).Scan(&currentPassword)
	if err != nil{
		return err
	}

	if currentPassword == hashedPassword{
		_, err = database.DB.Exec("UPDATE users SET password = ? WHERE id = ?", hashNewPassword, id)
	} else{
		return errors.New("Your old password doesn't mach that you have given!")
	}

	return nil
}

func GetUserId(username, password string) (string, error) {
	database.Connect()
	defer database.Close()

	var id string
	err := database.DB.QueryRow("SELECT id FROM users WHERE password = ? AND username = ?", password, username).Scan(&id)
	if err != nil{
		return "", err
	}

	return id ,nil
}

func GetUserPasswordAndUsername(id string) (string, string, error){
	database.Connect()
	defer database.Close()

	var password, username  string
	err := database.DB.QueryRow("SELECT password AND username FROM users WHERE id = ?", id).Scan(&password, &username)
	if err != nil{
		return "", "", err
	}

	return username, password, nil
}