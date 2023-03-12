package api

import (
	"strings"

	"github.com/Meonako/webui-api/sampler"
)

type Config struct {
	// URL to API endpoint (e.g. http://127.0.0.1:7860, https://b645a912.gradio.app)
	//
	// - Default: http://127.0.0.1:7860
	BaseURL string

	// API path is stored here
	Path *APIPath

	// Default Value are store here.
	Default *Default
}

type Default struct {
	// Sampling Method or Sampler (e.g. Euler a, DPM++ 2M Karras). You can type it in yourself or use built-in Helper Package: sampler
	//
	//   Default: "Euler a"
	Sampler string

	// Sampling Steps (e.g. 20, 120)
	//
	//   Default: 20
	Steps int

	// Classifier-Free Guidance Scale (e.g. 7, 12.0)
	//
	//   Default: 7.0
	CFGScale float64

	// Width of the image (e.g. 512, 1024)
	//
	//   Default: 512
	Width int

	// Height of the image (e.g. 512, 1024)
	//
	//   Default: 512
	Height int
}

type APIPath struct {
	// Path to txt2img API
	//
	//  - Default: /sdapi/v1/txt2img
	Txt2Img string

	// Path to img2img API
	//
	//  - Default: /sdapi/v1/img2img
	Img2Img string

	// Path to extra single image API
	//
	//  - Default: /sdapi/v1/extra-single-image
	ExtraSingle string

	// Path to extra batch images API
	//
	//  - Default: /sdapi/v1/extra-batch-images
	ExtraBatch string

	// Path to png info API
	//
	//  - Default: /sdapi/v1/png-info
	PNGInfo string

	// Path to progress API
	//
	//  - Default: /sdapi/v1/progress
	Progress string

	// Path to interrogate API
	//
	//  - Default: /sdapi/v1/interrogate
	Interrogate string

	// Path to interrupt API
	//
	//  - Default: /sdapi/v1/interrupt
	Interrupt string

	// Path to skip API
	//
	//  - Default: /sdapi/v1/skip
	Skip string

	// Path to options API
	//
	//  - Default: /sdapi/v1/options
	Options string

	// Path to sd-models API
	//
	//  - Default: /sdapi/v1/sd-models
	SDModels string
}

var DefaultConfig = Config{
	BaseURL: "http://127.0.0.1:7860",
	Path: &APIPath{
		Txt2Img:     "/sdapi/v1/txt2img",
		Img2Img:     "/sdapi/v1/img2img",
		ExtraSingle: "/sdapi/v1/extra-single-image",
		ExtraBatch:  "/sdapi/v1/extra-batch-images",
		PNGInfo:     "/sdapi/v1/png-info",
		Progress:    "/sdapi/v1/progress",
		Interrogate: "/sdapi/v1/interrogate",
		Interrupt:   "/sdapi/v1/interrupt",
		Skip:        "/sdapi/v1/skip",
		Options:     "/sdapi/v1/options",
		SDModels:    "/sdapi/v1/sd-models",
	},
}

/*
Default Values.

	Sampler  = sampler.EULER_A,
	Steps    = 28,
	CFGScale = 7,
	Width    = 512,
	Height   = 512,
*/
var DefaultValue = &Default{
	Sampler:  sampler.EULER_A,
	Steps:    28,
	CFGScale: 7,
	Width:    512,
	Height:   512,
}

func setDefault(conf ...Config) Config {
	if len(conf) <= 0 {
		return DefaultConfig
	}

	config := conf[0]

	if config.BaseURL == "" {
		config.BaseURL = DefaultConfig.BaseURL
	}

	config.BaseURL = strings.TrimSuffix(config.BaseURL, "/")

	if config.Path == nil {
		config.Path = DefaultConfig.Path
	}

	return config
}
