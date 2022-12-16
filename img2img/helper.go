package img2img

import (
	"bytes"
	"image"
	"image/png"
	"os"

	"github.com/Meonako/webui-api/utils"
)

// Convenient function to encode image in any format to base64
func ReadFromFile(path string) (string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return ReadFromBytes(fileBytes)
}

// Convenient function to encode image in any format to base64 and IGNORE ANY ERRORS that might occur.
func ReadFromFileIgnore(path string) string {
	fileBytes, _ := os.ReadFile(path)
	data, _ := ReadFromBytes(fileBytes)
	return data
}

// Convenient function to encode image data in bytes format to base64.
func ReadFromBytes(data []byte) (string, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	return "data:image/png;base64," + utils.EncodeBase64(buf.Bytes()), err
}
