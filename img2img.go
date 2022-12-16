package api

import (
	"github.com/goccy/go-json"
)

type Img2Img struct {
	InitImages             []string       `json:"init_images"`
	ResizeMode             int            `json:"resize_mode"`
	DenoisingStrength      float64        `json:"denoising_strength"`
	Mask                   string         `json:"mask"`
	MaskBlur               int            `json:"mask_blur"`
	InpaintingFill         int            `json:"inpainting_fill"`
	InpaintFullRes         bool           `json:"inpaint_full_res"`
	InpaintFullResPadding  int            `json:"inpaint_full_res_padding"`
	InpaintingMaskInvert   int            `json:"inpainting_mask_invert"`
	InitialNoiseMultiplier int            `json:"initial_noise_multiplier"`
	Prompt                 string         `json:"prompt"`
	Styles                 []string       `json:"styles"`
	Seed                   int            `json:"seed"`
	Subseed                int            `json:"subseed"`
	SubseedStrength        int            `json:"subseed_strength"`
	SeedResizeFromH        int            `json:"seed_resize_from_h"`
	SeedResizeFromW        int            `json:"seed_resize_from_w"`
	SamplerName            string         `json:"sampler_name"`
	BatchSize              int            `json:"batch_size"`
	BatchCount             int            `json:"n_iter"`
	Steps                  int            `json:"steps"`
	CFGScale               float64        `json:"cfg_scale"`
	Width                  int            `json:"width"`
	Height                 int            `json:"height"`
	RestoreFaces           bool           `json:"restore_faces"`
	Tiling                 bool           `json:"tiling"`
	NegativePrompt         string         `json:"negative_prompt"`
	Eta                    float64        `json:"eta"`
	SChurn                 float64        `json:"s_churn"`
	STmax                  float64        `json:"s_tmax"`
	STmin                  float64        `json:"s_tmin"`
	SNoise                 float64        `json:"s_noise"`
	OverrideSettings       map[string]any `json:"override_settings"`
	SamplerIndex           string         `json:"sampler_index"`
	IncludeInitImages      bool           `json:"include_init_images"`
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
