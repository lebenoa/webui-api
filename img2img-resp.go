package api

import (
	"bytes"
	"fmt"

	"github.com/Meonako/webui-api/utils"

	"github.com/goccy/go-json"
)

type img2imgRespond struct {
	Images        []string `json:"images"`
	DecodedImages [][]byte
	Parameters    Img2Img `json:"parameters"`
	Info          string  `json:"info"`
}

func newImg2ImgResp() *img2imgRespond {
	return &img2imgRespond{
		Images:        []string{},
		DecodedImages: [][]byte{},
	}
}

// Decode data at index store in "Images" field and return it.
func (tr *img2imgRespond) DecodeImage(index int) ([]byte, error) {
	if len(tr.Images) <= index {
		return []byte{}, fmt.Errorf("%v", "Out of bound. Provided Index > len(Images) of struct.")
	}

	return utils.DecodeBase64(tr.Images[index])
}

// Decode all data store in "Images" field and return it. You can access this later in "DecodedImages" Field.
func (tr *img2imgRespond) DecodeAllImages() ([][]byte, error) {
	if tr.DecodedImages == nil || len(tr.DecodedImages) > 0 {
		tr.DecodedImages = [][]byte{}
	}

	for index := range tr.Images {
		imageData, err := tr.DecodeImage(index)
		if err != nil {
			return [][]byte{}, err
		}

		tr.DecodedImages = append(tr.DecodedImages, imageData)
	}
	return tr.DecodedImages, nil
}

// Make bytes.Reader from "DecodedImages" field.
//
//   - It'll call "DecodeAllImages()" if "len(DecodedImages) <= 0" then continue to proceed.
//
// This is ready to send to discord. Or ready to png.Decode and save.
func (tr *img2imgRespond) MakeBytesReader() (reader []*bytes.Reader, err error) {
	if tr.DecodedImages == nil || len(tr.DecodedImages) > 0 {
		tr.DecodedImages = [][]byte{}
	}

	for index := range tr.Images {
		imageData, err := tr.DecodeImage(index)
		if err != nil {
			return nil, err
		}

		reader = append(reader, bytes.NewReader(imageData))
		tr.DecodedImages = append(tr.DecodedImages, imageData)
	}

	return reader, nil
}

// Info field contains generation parameters like "Parameters" field but in long string instead.
//   - So I wouldn't recommend doing this as it intend to be use as long string *UNLESS* you know what you're doing.
func (tr *img2imgRespond) DecodeInfo() (res map[string]any, err error) {
	err = json.Unmarshal([]byte(tr.Info), &res)
	return
}

// Upscale any images in "Images" field.
//
//	Leave `params.ImagesList` field empty to upscale all images.
func (tr *img2imgRespond) Upscale(params *ExtraBatchImages) (*extraBatchImagesRespond, error) {
	if params.ImagesList == nil || len(params.ImagesList) <= 0 {
		params.ImagesList = tr.buildBatch()
	}

	return API.ExtraBatchImages(params)
}

func (tr *img2imgRespond) buildBatch() (res []ImageData) {
	for _, image := range tr.Images {
		res = append(res, ImageData{Data: "data:image/png;base64," + image})
	}
	return
}
