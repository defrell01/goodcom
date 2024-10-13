package files

import (
	"os"
	"path/filepath"
	"strings"
)

func ScanDirectory(root string, extensions []string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && hasValidExtension(path, extensions) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func hasValidExtension(filePath string, extensions []string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(filePath, ext) {
			return true
		}
	}
	return false
}
