package utils

import (
	"os"
	"encoding/base64"
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
	return EncodeBase64(fileBytes), err
}

// Convenience function to encode image in any format to base64 and IGNORE ANY ERRORS that might occur.
func Base64FromFileIgnore(path string) string {
	fileBytes, _ := os.ReadFile(path)
	return EncodeBase64(fileBytes)
}