package api

import (
	"fmt"

	"github.com/goccy/go-json"
)

type Txt2Image struct {
	EnableHR          bool           `json:"enable_hr,omitempty"` // Hi-res fix
	DenoisingStrength float64        `json:"denoising_strength,omitempty"`
	FirstphaseWidth   int            `json:"firstphase_width,omitempty"`
	FirstphaseHeight  int            `json:"firstphase_height,omitempty"`
	Prompt            string         `json:"prompt"`
	NegativePrompt    string         `json:"negative_prompt,omitempty"`
	Styles            []string       `json:"styles,omitempty"`
	Seed              int64          `json:"seed,omitempty"`
	Subseed           int            `json:"subseed,omitempty"`
	SubseedStrength   int            `json:"subseed_strength,omitempty"`
	SeedResizeFromH   int            `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW   int            `json:"seed_resize_from_w,omitempty"`
	SamplerName       string         `json:"sampler_name,omitempty"`
	SamplerIndex      string         `json:"sampler_index,omitempty"`
	BatchSize         int            `json:"batch_size,omitempty"`
	BatchCount        int            `json:"n_iter,omitempty"`
	Steps             int            `json:"steps,omitempty"`
	CFGScale          float64        `json:"cfg_scale,omitempty"`
	Width             int            `json:"width,omitempty"`
	Height            int            `json:"height,omitempty"`
	RestoreFaces      bool           `json:"restore_faces,omitempty"`
	Tiling            bool           `json:"tiling,omitempty"`
	Eta               float64        `json:"eta,omitempty"`
	SChurn            float64        `json:"s_churn,omitempty"`
	STmax             int            `json:"s_tmax,omitempty"`
	STmin             float64        `json:"s_tmin,omitempty"`
	SNoise            float64        `json:"s_noise,omitempty"`
	OverrideSettings  map[string]any `json:"override_settings,omitempty"`
}

func (data *Txt2Image) processDefault(a *api) {
	if !a.Config.UseDefault {
		return
	}

	if data.Width == 0 {
		data.Width = a.Config.DefaultWidth
	} else if data.Width%64 != 0 {
		data.Width = data.Width - (data.Width % 64)
	}

	if data.Height == 0 {
		data.Height = a.Config.DefaultHeight
	} else if data.Height%64 != 0 {
		data.Height = data.Height - (data.Height % 64)
	}

	if data.CFGScale == 0 {
		data.CFGScale = a.Config.DefaultCFGScale
	}

	if data.Steps == 0 {
		data.Steps = a.Config.DefaultSteps
	}

	if data.SamplerName == "" {
		data.SamplerName = a.Config.DefaultSampler
	}

	if data.SamplerIndex == "" {
		data.SamplerIndex = a.Config.DefaultSampler
	}
}

// Generate Image based on Text. Return Respond struct and Error object.
func (a *api) Text2Image(params Txt2Image) (res *txt2ImageRespond, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recover from Txt2Img: %v", r)
		}
	}()

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
