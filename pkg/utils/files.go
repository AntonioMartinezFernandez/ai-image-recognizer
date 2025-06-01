package utils

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func JpgFileNames(folder string) ([]string, error) {
	var jpgFiles []string

	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".jpg") {
			jpgFiles = append(jpgFiles, d.Name())
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	return jpgFiles, nil
}
