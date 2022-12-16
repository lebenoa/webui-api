package img2img

import (
	"bytes"
	"image"
	"image/png"
	"os"

	"github.com/Meonako/webui-api/utils"
)

func ReadFromFile(path string) (string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return ReadFromBytes(fileBytes)
}

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

	return utils.EncodeBase64(buf.Bytes()), err
}
