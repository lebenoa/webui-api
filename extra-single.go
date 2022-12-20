package api

import "github.com/goccy/go-json"

type ExtraSingleImage struct {
	ResizeMode                int     `json:"resize_mode,omitempty"`
	ShowExtrasResults         bool    `json:"show_extras_results,omitempty"`
	GfpganVisibility          float64 `json:"gfpgan_visibility,omitempty"`
	CodeformerVisibility      float64 `json:"codeformer_visibility,omitempty"`
	CodeformerWeight          float64 `json:"codeformer_weight,omitempty"`
	UpscalingResize           int     `json:"upscaling_resize,omitempty"`
	UpscalingResizeW          int     `json:"upscaling_resize_w,omitempty"`
	UpscalingResizeH          int     `json:"upscaling_resize_h,omitempty"`
	UpscalingCrop             bool    `json:"upscaling_crop,omitempty"`
	Upscaler1                 string  `json:"upscaler_1,omitempty"`
	Upscaler2                 string  `json:"upscaler_2,omitempty"`
	ExtrasUpscaler2Visibility float64 `json:"extras_upscaler_2_visibility,omitempty"`
	UpscaleFirst              bool    `json:"upscale_first,omitempty"`
	Image                     string  `json:"image,omitempty"`
}

func (a *api) ExtraSingleImage(params ExtraSingleImage) (*extraSingleImageRespond, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return &extraSingleImageRespond{}, err
	}

	data, err := a.post(a.Config.Path.ExtraSingle, payload)
	if err != nil {
		return &extraSingleImageRespond{}, err
	}

	apiResp := extraSingleImageRespond{}
	err = json.Unmarshal(data, &apiResp)
	return &apiResp, err
}
