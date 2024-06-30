package util

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
)

const uploadPath = "../markdown"

type LocalFileManager struct {
}

func (lfm LocalFileManager) SaveFile(file multipart.File, filename string) (string, error) {
	localFilename := uploadPath + "/" + filename // TODO: Create a unique file name
	out, err := os.Create(localFilename)
	if err != nil {
		log.Printf("error creating file")
		return "", errors.New("error creating file")
	}
	_, err = io.Copy(out, file)
	if err != nil {
		return "", errors.New("error populating file")
	}

	return localFilename, nil
}

func (lfm LocalFileManager) GetFile(filename string) ([]byte, error) {
	return os.ReadFile(uploadPath + filename + ".md")
}
