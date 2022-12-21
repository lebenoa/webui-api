package utils

import (
	"encoding/base64"
	"os"
)

func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func DecodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// Convenience function to encode image in any format to base64
func Base64FromFile(path string) (string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return EncodeBase64(fileBytes), nil
}

// Convenience function to encode image in any format to base64 and IGNORE ANY ERRORS that might occur.
func Base64FromFileIgnore(path string) string {
	fileBytes, _ := os.ReadFile(path)
	return EncodeBase64(fileBytes)
}

func Base64FromFiles(allPath ...string) (res []string, err error) {
	for _, path := range allPath {
		encoded, err := Base64FromFile(path)
		if err != nil {
			return []string{}, err
		}
		res = append(res, encoded)
	}

	return
}
