package api

import (
	"github.com/Meonako/webui-api/utils"
	"github.com/goccy/go-json"
)

type ExtraBatchImages struct {
	// Resize Mode: Scale By || Scale To. See: webup-api/extra package for helper.
	//
	//  Default: 0 (Scale By)
	ResizeMode int `json:"resize_mode,omitempty"`

	// Don't even know what this is.
	//
	//  Default: true
	ShowExtrasResults bool `json:"show_extras_results,omitempty"`

	// GFPGAN Face restoration. Value must be between 0.0 - 1.0
	//
	//  Default: 0.0
	GfpganVisibility float64 `json:"gfpgan_visibility,omitempty"`

	// CodeFormer Face restoration. Value must be between 0.0 - 1.0
	//
	//  Default: 0.0
	CodeformerVisibility float64 `json:"codeformer_visibility,omitempty"`

	// CodeFormer Face restoration weight. 0 = Maximum Effect, 1 = Minimum Effect.
	//
	//  Default: 0.0
	CodeformerWeight float64 `json:"codeformer_weight,omitempty"`

	// Multiplier to width and height of the original image.
	//
	//  NOTE: Will only work if ResizeMode is 0
	//  Default: 2
	UpscalingResize int `json:"upscaling_resize,omitempty"`

	// Width of the result image.
	//
	//  NOTE: Will only work if ResizeMode is 1
	//  Default: 512
	UpscalingResizeW int `json:"upscaling_resize_w,omitempty"`

	// Height of the result image.
	//
	//  NOTE: Will only work if ResizeMode is 1
	//  Default: 512
	UpscalingResizeH int `json:"upscaling_resize_h,omitempty"`

	// Crop Image if the aspect ratio of original image and result image doesn't match.
	//
	//  NOTE: Will only work if ResizeMode is 1
	//  Default: true
	UpscalingCrop bool `json:"upscaling_crop,omitempty"`

	// First Upscaler Model. See: webui-api/extra package for helper.
	//
	//  Default: "None"
	Upscaler1 string `json:"upscaler_1,omitempty"`

	// Second Upscaler Model. See: webui-api/extra package for helper.
	//
	//  Default: "None"
	Upscaler2 string `json:"upscaler_2,omitempty"`

	// Second Upscaler Model Visibility. See: webui-api/extra package for helper.
	//
	//  Default: 0.0
	ExtrasUpscaler2Visibility float64 `json:"extras_upscaler_2_visibility,omitempty"`

	// Upscale first then do face restoration.
	//
	//  Default: false
	UpscaleFirst bool `json:"upscale_first,omitempty"`

	// Base64-encoded image to be upscale.
	//
	//  Default: ""
	ImagesList []ImageData `json:"imageList,omitempty"`
}

type ImageData struct {
	// Base64-encoded image to be upscale.
	//
	//  Default: ""
	Data string `json:"data"`

	// I don't know what this is. I tried to read the source code what it does but I don't think I get it.
	//
	// **NOT CONFIRM** Perhaps it is the temp file name
	//  Default: ""
	Name string `json:"name,omitempty"`
}

// Convenience function to build []ImageData from base64-encoded image
func BuildBatch(imageList ...string) []ImageData {
	imageData := []ImageData{}

	for _, image := range imageList {
		imageData = append(imageData, ImageData{Data: "data:image/png;base64," + image})
	}

	return imageData
}

// Convenience function to build []ImageData from files
func BuildBatchFromFiles(files ...string) ([]ImageData, error) {
	imageData := []ImageData{}

	for _, file := range files {
		b64Data, err := utils.Base64FromFile(file)
		if err != nil {
			return []ImageData{}, err
		}
		imageData = append(imageData, ImageData{Data: "data:image/png;base64," + b64Data})
	}

	return imageData, nil
}

func (a *api) ExtraBatchImages(params *ExtraBatchImages, decode ...bool) (*extraBatchImagesRespond, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return &extraBatchImagesRespond{}, err
	}

	data, err := a.post(a.Config.Path.ExtraBatch, payload)
	if err != nil {
		return &extraBatchImagesRespond{}, err
	}

	apiResp := extraBatchImagesRespond{}
	err = json.Unmarshal(data, &apiResp)

	if len(decode) <= 0 {
		return &apiResp, err
	}

	apiResp.DecodeAllImages()
	return &apiResp, err
}