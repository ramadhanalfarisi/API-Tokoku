package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)


func Uploader(uploadedFile interface{}, handler interface{}, alias string) error{
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	filename := handler.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return err
	}

	return nil
}
