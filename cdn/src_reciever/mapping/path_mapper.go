package mapping

import (
	"fmt"
	"os"
)

func EnsureImagePath(uid string, username string, imageType string) (string, error) {
	basePath := "src_reciever/storage"
	fullPath := fmt.Sprintf("%s/%s/%s/%s", basePath, uid, username, imageType)

	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}
