package api

import (
	"github.com/goccy/go-json"
)

type Txt2Image struct {
	EnableHR          bool    `json:"enable_hr,omitempty"`          // Hi-res fix.
	DenoisingStrength float64 `json:"denoising_strength,omitempty"` // Hi-res fix option. Determines how little respect the algorithm should have for image's content. At 0, nothing will change, and at 1 you'll get an unrelated image.
	FirstphaseWidth   int     `json:"firstphase_width,omitempty"`   // Hi-res fix option. Might not work anymore
	FirstphaseHeight  int     `json:"firstphase_height,omitempty"`  // Hi-res fix option. Might not work anymore

	// Hi-res fix option. Multiplier to original width and height.
	//
	// HRScale = 2 will work like this: 384x512 will result in 768x1024
	//
	//  Only HRScale or HRResizeX / HRResizeY will be used
	HRScale float64 `json:"hr_scale,omitempty"`

	// Hi-res fix option. Which Hi-res upscale model will be used.
	//
	//  See: `upscaler` helper package (github.com/Meonako/webui-api/upscaler)
	HRUpscaler string `json:"hr_upscaler,omitempty"`

	// Hi-res fix option. After denoising and upscale, use this amount of steps instead of the amount before denoise and upscale.
	HRSecondPassSteps int `json:"hr_second_pass_steps,omitempty"`

	// Hi-res fix option. The width of the result image
	//
	//  Only HRScale or HRResizeX / HRResizeY will be used
	HRResizeX int `json:"hr_resize_x,omitempty"`

	// Hi-res fix option. The height of the result image
	//
	//  Only HRScale or HRResizeX / HRResizeY will be used
	HRResizeY int `json:"hr_resize_y,omitempty"`

	Prompt           string         `json:"prompt"`
	NegativePrompt   string         `json:"negative_prompt,omitempty"`
	Styles           []string       `json:"styles,omitempty"`
	Seed             int64          `json:"seed,omitempty"` // A value that determines the output of random number generator - if you create an image with same parameters and seed as another image, you'll get the same result
	Subseed          int            `json:"subseed,omitempty"`
	SubseedStrength  int            `json:"subseed_strength,omitempty"`
	SeedResizeFromH  int            `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW  int            `json:"seed_resize_from_w,omitempty"`
	SamplerName      string         `json:"sampler_name,omitempty"`  // Either SamplerName or SamplerIndex will be used.
	SamplerIndex     string         `json:"sampler_index,omitempty"` // Either SamplerName or SamplerIndex will be used.
	BatchSize        int            `json:"batch_size,omitempty"`    // How many do you want to simultaneously generate.
	BatchCount       int            `json:"n_iter,omitempty"`        // How many times do you want to generate.
	Steps            int            `json:"steps,omitempty"`         // How many times to improve the generated image iteratively; higher values take longer; very low values can produce bad results
	CFGScale         float64        `json:"cfg_scale,omitempty"`     // Classifier Free Guidance Scale - how strongly the image should conform to prompt - lower values produce more creative results
	Width            int            `json:"width,omitempty"`
	Height           int            `json:"height,omitempty"`
	RestoreFaces     bool           `json:"restore_faces,omitempty"`
	Tiling           bool           `json:"tiling,omitempty"`
	DoNotSaveSamples bool           `json:"do_not_save_samples,omitempty"`
	DoNotSaveGrid    bool           `json:"do_not_save_grid,omitempty"`
	Eta              float64        `json:"eta,omitempty"`
	SChurn           float64        `json:"s_churn,omitempty"`
	STmax            int            `json:"s_tmax,omitempty"`
	STmin            float64        `json:"s_tmin,omitempty"`
	SNoise           float64        `json:"s_noise,omitempty"`
	OverrideSettings map[string]any `json:"override_settings,omitempty"`

	// Since API default value of this field is `true`, it can't be `omitempty`.
	//
	//  That means the default value of this lib is false
	OverrideSettingsRestoreAfterwards bool `json:"override_settings_restore_afterwards"`

	ScriptName string `json:"script_name,omitempty"`
	ScriptArgs string `json:"script_args,omitempty"`

	// Since API default value of this field is `true`, it can't be `omitempty`.
	//
	//  That means the default value of this lib is false
	SendImages bool `json:"send_images"`

	SaveImages bool `json:"save_iamges,omitempty"`
}

func (data *Txt2Image) processDefault(a *api) {
	if a.Config.Default == nil {
		return
	}

	if data.Width == 0 {
		data.Width = a.Config.Default.Width
	}

	if data.Height == 0 {
		data.Height = a.Config.Default.Height
	}

	if data.CFGScale == 0 {
		data.CFGScale = a.Config.Default.CFGScale
	}

	if data.Steps == 0 {
		data.Steps = a.Config.Default.Steps
	}

	if data.SamplerName == "" {
		data.SamplerName = a.Config.Default.Sampler
	}

	if data.SamplerIndex == "" {
		data.SamplerIndex = a.Config.Default.Sampler
	}
}

// Generate Image based on Text. Return Respond struct and Error object.
func (a *api) Text2Image(params *Txt2Image) (*txt2ImageRespond, error) {
	params.processDefault(a)

	payload, err := json.Marshal(params)
	if err != nil {
		return &txt2ImageRespond{}, err
	}

	data, err := a.post(a.Config.Path.Txt2Img, payload)
	if err != nil {
		return &txt2ImageRespond{}, err
	}

	apiResp := newTxt2ImgResp()
	err = json.Unmarshal(data, &apiResp)
	return apiResp, err
}
