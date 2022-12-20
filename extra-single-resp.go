package api

import "github.com/Meonako/webui-api/utils"

type extraSingleImageRespond struct {
	HTMLInfo     string `json:"html_info"` // Upscaler info in HTML format. I don't even know why they return HTML format for API
	Image        string `json:"image"`     // Base64-encoded image
	DecodedImage []byte // Base64-decoded image data in byte
}

func (er *extraSingleImageRespond) DecodeImage() (decoded []byte, err error) {
	decoded, err = utils.DecodeBase64(er.Image)
	er.DecodedImage = decoded
	return
}
