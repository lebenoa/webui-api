package api

import (
	"bytes"
	"fmt"

	"github.com/Meonako/webui-api/utils"

	"github.com/goccy/go-json"
)

type txt2ImageRespond struct {
	Images        []string  `json:"images"`     // Base64-encoded Images Data
	DecodedImages [][]byte  `json:"-"`          // Base64-decoded Images Data store here after "DecodeAllImages()" called
	Parameters    Txt2Image `json:"parameters"` // Generation Parameters. Should be the same value as the one you pass to generate.
	Info          string    `json:"info"`       // Info field contains generation parameters like "parameters" field but in long string instead.
}

func newTxt2ImgResp() *txt2ImageRespond {
	return &txt2ImageRespond{
		Images:        []string{},
		DecodedImages: [][]byte{},
	}
}

// Decode data at index store in "Images" field and return it.
func (tr *txt2ImageRespond) DecodeImage(index int) ([]byte, error) {
	if len(tr.Images) <= index {
		return []byte{}, fmt.Errorf("%v", "Out of bound. Provided Index > len(Images) of struct.")
	}

	return utils.DecodeBase64(tr.Images[index])
}

// Decode all data store in "Images" field and return it. You can access this later in "DecodedImages" Field.
func (tr *txt2ImageRespond) DecodeAllImages() ([][]byte, error) {
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
func (tr *txt2ImageRespond) MakeBytesReader() (reader []*bytes.Reader, err error) {
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
func (tr *txt2ImageRespond) DecodeInfo() (res map[string]any, err error) {
	err = json.Unmarshal([]byte(tr.Info), &res)
	return
}
