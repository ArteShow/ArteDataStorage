package repository

import (
	"errors"
	"io"
	"os"
)

const tempFolderPath string = "./internal/temp/"
const fileFolderPath string = "./internal/files/"

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isFolderEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)

	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func CreateUserFolder(userID string) error {
	exists := folderExists(fileFolderPath + userID)
	if exists {
		return errors.New("user folder already exists")
	}

	if err := os.MkdirAll(fileFolderPath+userID, 0755); err != nil {
		return err
	}

	return nil
}

func DeleteTempFile(name, userID string) error {
	exists := fileExists(tempFolderPath + userID + "/" + name)
	if !exists {
		return errors.New("temp file not found")
	}

	if err := os.Remove(tempFolderPath + userID + "/" + name); err != nil {
		return err
	}

	empty, err := isFolderEmpty(tempFolderPath + userID)
	if err != nil {
		return err
	}

	if empty {
		if err := os.RemoveAll(tempFolderPath + userID); err != nil {
			return err
		}
	}

	return nil
}

func CreateUserTempFolder(userID string) error {
	exists := folderExists(tempFolderPath + userID)
	if exists {
		return errors.New("user temp folder already exists")
	}

	if err := os.Mkdir(tempFolderPath+userID, 0755); err != nil {
		return err
	}

	return nil
}

func MoveFile(userID, path, name string) error {
	exists := fileExists(tempFolderPath + userID + "/" + name)
	if !exists {
		return errors.New("temp file not found")
	}

	exists = folderExists(fileFolderPath + userID + "/" + path)
	if !exists {
		return errors.New("destination folder not found")
	}

	if err := copyFile(tempFolderPath+userID+"/"+name, fileFolderPath+userID+"/"+path); err != nil {
		return err
	}

	if err := DeleteTempFile(name, userID); err != nil {
		return err
	}

	return nil
}

func DeleteUserFile(userID, name, path string) error {
	exists := fileExists(fileFolderPath + userID + "/" + path + "/" + name)
	if !exists {
		return errors.New("file not found under this path")
	}

	if err := os.Remove(fileFolderPath + userID + "/" + path + "/" + name); err != nil {
		return err
	}

	return nil
}
