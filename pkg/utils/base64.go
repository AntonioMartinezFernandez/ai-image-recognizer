package utils

import (
	"encoding/base64"
	"fmt"
	"os"
)

func EncodeJpgBase64(filePath string) (*string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	dataURI := "data:image/jpeg;base64," + base64Data

	return &dataURI, nil
}
