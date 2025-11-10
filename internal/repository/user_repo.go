package repository

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/ArteShow/ArteDataStorage/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) (bool, error, string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err, ""
	}

	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return false, err, ""
	}

	id := hex.EncodeToString(bytes)

	database.Connect()
	defer database.Close()

	database.Migrate()

	_, err2 := database.DB.Exec(
		"INSERT INTO users (id, username, password) VALUES (?, ?, ?)",
		id, username, hash,
	)

	if err2 != nil {
		return false, err2, ""
	}

	return true, nil, id
}
