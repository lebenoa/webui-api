package api

import (
	"github.com/goccy/go-json"
)

type Img2Img struct {
	InitImages             []string       `json:"init_images,omitempty"`
	ResizeMode             int            `json:"resize_mode,omitempty"`
	DenoisingStrength      float64        `json:"denoising_strength,omitempty"`
	Mask                   string         `json:"mask,omitempty"`
	MaskBlur               int            `json:"mask_blur,omitempty"`
	InpaintingFill         int            `json:"inpainting_fill,omitempty"`
	InpaintFullRes         bool           `json:"inpaint_full_res,omitempty"`
	InpaintFullResPadding  int            `json:"inpaint_full_res_padding,omitempty"`
	InpaintingMaskInvert   int            `json:"inpainting_mask_invert,omitempty"`
	InitialNoiseMultiplier int            `json:"initial_noise_multiplier,omitempty"`
	Prompt                 string         `json:"prompt,omitempty"`
	Styles                 []string       `json:"styles,omitempty"`
	Seed                   int            `json:"seed,omitempty"`
	Subseed                int            `json:"subseed,omitempty"`
	SubseedStrength        int            `json:"subseed_strength,omitempty"`
	SeedResizeFromH        int            `json:"seed_resize_from_h,omitempty"`
	SeedResizeFromW        int            `json:"seed_resize_from_w,omitempty"`
	SamplerName            string         `json:"sampler_name,omitempty"`
	BatchSize              int            `json:"batch_size,omitempty"`
	BatchCount             int            `json:"n_iter,omitempty"`
	Steps                  int            `json:"steps,omitempty"`
	CFGScale               float64        `json:"cfg_scale,omitempty"`
	Width                  int            `json:"width,omitempty"`
	Height                 int            `json:"height,omitempty"`
	RestoreFaces           bool           `json:"restore_faces,omitempty"`
	Tiling                 bool           `json:"tiling,omitempty"`
	NegativePrompt         string         `json:"negative_prompt,omitempty"`
	Eta                    float64        `json:"eta,omitempty"`
	SChurn                 float64        `json:"s_churn,omitempty"`
	STmax                  float64        `json:"s_tmax,omitempty"`
	STmin                  float64        `json:"s_tmin,omitempty"`
	SNoise                 float64        `json:"s_noise,omitempty"`
	OverrideSettings       map[string]any `json:"override_settings,omitempty"`
	SamplerIndex           string         `json:"sampler_index,omitempty"`
	IncludeInitImages      bool           `json:"include_init_images,omitempty"`
}

func (i *Img2Img) processDefault(a *api) {
	if !a.Config.UseDefault {
		return
	}

	if i.Width == 0 {
		i.Width = a.Config.DefaultWidth
	} else if i.Width%64 != 0 {
		i.Width = i.Width - (i.Width % 64)
	}

	if i.Height == 0 {
		i.Height = a.Config.DefaultHeight
	} else if i.Height%64 != 0 {
		i.Height = i.Height - (i.Height % 64)
	}

	if i.CFGScale == 0 {
		i.CFGScale = a.Config.DefaultCFGScale
	}

	if i.Steps == 0 {
		i.Steps = a.Config.DefaultSteps
	}

	if i.SamplerName == "" {
		i.SamplerName = a.Config.DefaultSampler
	}

	if i.SamplerIndex == "" {
		i.SamplerIndex = a.Config.DefaultSampler
	}
}

func (a *api) Image2Image(i Img2Img) (*img2imgRespond, error) {
	i.processDefault(a)

	payload, err := json.Marshal(i)
	if err != nil {
		return &img2imgRespond{}, err
	}

	data, err := a.post(a.Config.Path.Txt2Img, payload)
	if err != nil {
		return &img2imgRespond{}, err
	}

	apiResp := newImg2ImgResp()
	err = json.Unmarshal(data, &apiResp)
	return apiResp, err
}
