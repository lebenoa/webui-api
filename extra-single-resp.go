package api

import "github.com/Meonako/webui-api/utils"

type extraSingleImageRespond struct {
	HTMLInfo     string `json:"html_info"` // Upscaler info in HTML format.
	Image        string `json:"image"`     // Base64-encoded image
	DecodedImage []byte // Base64-decoded image data in byte
}

func (er *extraSingleImageRespond) DecodeImage() ([]byte, error) {
	return utils.DecodeBase64(er.Image)
}
