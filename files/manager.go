package files

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalFileManager struct {
	UploadPath string
}

func NewLocalFileManager(uploadPath string) LocalFileManager {
	return LocalFileManager{
		UploadPath: uploadPath,
	}
}

// Saves file to the upload path set in the configuration. If the file exists it overwrites it
func (lfm LocalFileManager) SaveFile(file multipart.File, filename string) (string, error) {
	log.Println("saving file")
	localFilename := filepath.Join(lfm.UploadPath, filename)

	err := os.MkdirAll(lfm.UploadPath, os.ModePerm)
	if err != nil {
		log.Printf("error creating upload directory: %v", err)
		return "", errors.New("error creating directory")
	}

	out, err := os.Create(localFilename)
	if err != nil {
		log.Printf("error creating file: %v", err)
		return "", errors.New("error creating file")
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", errors.New("error populating file")
	}

	return localFilename, nil
}

// Reads the file from the path set in the configuration.
func (lfm LocalFileManager) GetFile(filename string) ([]byte, error) {
	return os.ReadFile(lfm.UploadPath + filename + ".md")
}
