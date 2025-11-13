package repository

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/ArteShow/ArteDataStorage/internal/database"
)

func CreateFile(Name, UserId string, Type string) (string, error) {
	database.Connect()
	defer database.Close()

	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	id := hex.EncodeToString(bytes)

	_, err := database.DB.Exec("INSERT INTO files (userId, public, type, name, entryId) VALUES (?, ?, ?, ?, ?)", UserId, false, Type, Name, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func DeleteFile(UserId, EntryId string) error {
	database.Connect()
	defer database.Close()

	_, err := database.DB.Exec("DELETE FROM files WHERE userId = ? AND entryId = ?", UserId, EntryId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFileName(UserId, EntryId, NewName string) error {
	database.Connect()
	defer database.Close()

	_, err := database.DB.Exec("UPDATE files SET name = ? WHERE userId = ? AND entryId = ?", NewName, UserId, EntryId)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFileStatus(UserId, EntryId string, status bool) error {
	database.Connect()
	defer database.Close()

	_, err := database.DB.Exec("UPDATE files SET public = ? WHERE userId = ? AND entryId = ?", status, UserId, EntryId)
	if err != nil {
		return err
	}

	return err
}

func GetFileProperties(UserId, EntryId string) (string, string, bool, error) {
	database.Connect()
	defer database.Close()

	var FileName, FileType string
	var Status bool

	err := database.DB.QueryRow("SELECT name, public, type FROM files WHERE userId = ? AND entryId = ?", UserId, EntryId).Scan(&FileName, &Status, &FileType)
	if err != nil {
		return "", "", false, err
	}

	return FileName, FileType, Status, nil
}
