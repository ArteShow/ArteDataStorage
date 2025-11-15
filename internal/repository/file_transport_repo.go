package repository

import (
	"errors"
	"io"
	"os"
)

const TempFolderPath string = "./internal/temp/"
const FileFolderPath string = "./internal/files/"

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsFolderEmpty(path string) (bool, error) {
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

func CopyFile(src, dst string) error {
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

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func CreateUserFolder(userID string) (bool, error) {
	exists := FolderExists(FileFolderPath + userID)
	if exists {
		return false, errors.New("folder for this user already exists")
	}

	err := os.MkdirAll(FileFolderPath+userID, 0755)
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteTempFile(name, userID string) (bool, error) {
	exists := FileExists(TempFolderPath + userID + "/" + name)
	if !exists {
		return false, errors.New("there is no file with this name")
	}

	err := os.Remove(TempFolderPath + userID + "/" + name)
	if err != nil {
		return false, err
	}

	empty, err := IsFolderEmpty(TempFolderPath + userID)
	if err != nil {
		return false, err
	}

	if empty {
		err := os.RemoveAll(TempFolderPath + userID)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func CreateUserTempFolder(userID string) (bool, error) {
	exists := FolderExists(TempFolderPath + userID)
	if exists {
		return false, errors.New("folder for this user already exists")
	}

	err := os.Mkdir(TempFolderPath+userID, 0755)
	if err != nil {
		return false, err
	}

	return true, nil
}

func MoveFile(userID, path, name string) (bool, error) {
	exists := FileExists(TempFolderPath + userID + "/" + name)
	if !exists {
		return false, errors.New("file with this name does not exist")
	}

	exists = FolderExists(FileFolderPath + userID + "/" + path)
	if !exists {
		return false, errors.New("folder does not exist")
	}

	err := CopyFile(TempFolderPath+userID+"/"+name, FileFolderPath+userID+"/"+path)
	if err != nil {
		return false, err
	}

	if ok, err := DeleteTempFile(name, userID); err != nil || !ok {
		return false, err
	}

	return true, nil
}
