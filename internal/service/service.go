package service

import (
	"github.com/ArteShow/ArteDataStorage/internal/repository"
	"github.com/ArteShow/ArteDataStorage/pkg/logger"
)

func CreateFile(name, userID, formate, path string) (string, error) {
	DbEntryId, err := repository.CreateFile(name, userID, formate)
	if err != nil {
		logger.Log.Println("service.go: Failed to create an entry in the file db ", err)
		return "", err
	}

	err = repository.MoveFile(userID, path, name)
	if err != nil {
		logger.Log.Println("service.go: Failed to move the file ", err)
		return "", err
	}

	return DbEntryId, nil
}

func DeleteFile(name, userId, path, dbEntryId string) error {
	err := repository.DeleteFile(userId, dbEntryId)
	if err != nil {
		logger.Log.Println("service.go: Failed to delete the file ", err)
		return err
	}

	err = repository.DeleteUserFile(userId, name, path)
	if err != nil {
		logger.Log.Println("service.go: Failed to delete the file ", err)
		return err
	}

	return nil
}
