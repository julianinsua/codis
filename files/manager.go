package files

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
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
	localFilename := lfm.UploadPath + "/" + filename
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

// Reads the file from the path set in the configuration.
func (lfm LocalFileManager) GetFile(filename string) ([]byte, error) {
	return os.ReadFile(lfm.UploadPath + filename + ".md")
}
