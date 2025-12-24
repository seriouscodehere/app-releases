package mapping

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateFileName(dir string, baseName string, ext string) string {
	filename := fmt.Sprintf("%s%s", baseName, ext)
	fullPath := filepath.Join(dir, filename)

	counter := 1
	for {
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return filename
		}
		filename = fmt.Sprintf("%s%d%s", baseName, counter, ext)
		fullPath = filepath.Join(dir, filename)
		counter++
	}
}
