package api

import (
	"bytes"
	"fmt"

	"github.com/Meonako/webui-api/utils"
)

type extraBatchImagesRespond struct {
	HTMLInfo      string   `json:"html_info"` // Upscaler info in HTML format. I don't even know why they return HTML format for API
	Images        []string `json:"images"`    // Base64-encoded image
	DecodedImages [][]byte // Base64-decoded image data in byte
}

// Decode data at index store in "Images" field and return it.
func (er *extraBatchImagesRespond) DecodeImage(index int) ([]byte, error) {
	if len(er.Images) <= index {
		return []byte{}, fmt.Errorf("%v", "Out of bound. Provided Index > len(Images) of struct.")
	}

	return utils.DecodeBase64(er.Images[index])
}

// Decode all data store in "Images" field and return it. You can access this later in "DecodedImages" Field.
func (er *extraBatchImagesRespond) DecodeAllImages() ([][]byte, error) {
	if er.DecodedImages == nil || len(er.DecodedImages) > 0 {
		er.DecodedImages = [][]byte{}
	}

	for index := range er.Images {
		imageData, err := er.DecodeImage(index)
		if err != nil {
			return [][]byte{}, err
		}

		er.DecodedImages = append(er.DecodedImages, imageData)
	}

	return er.DecodedImages, nil
}

// Make bytes.Reader from "DecodedImages" field.
//
//   - It'll call "DecodeAllImages()" if "len(DecodedImages) <= 0" then continue to proceed.
//
// This is ready to send to discord. Or ready to png.Decode and save.
func (er *extraBatchImagesRespond) MakeBytesReader() (reader []*bytes.Reader, err error) {
	if er.DecodedImages == nil || len(er.DecodedImages) > 0 {
		er.DecodedImages = [][]byte{}
	}

	for index := range er.Images {
		imageData, err := er.DecodeImage(index)
		if err != nil {
			return nil, err
		}

		reader = append(reader, bytes.NewReader(imageData))
		er.DecodedImages = append(er.DecodedImages, imageData)
	}

	return reader, nil
}
