package api

import "github.com/Meonako/webui-api/sampler"

type Config struct {
	// URL to API endpoint (e.g. http://127.0.0.1:7860, https://b645a912.gradio.app)
	//
	// - Default: http://127.0.0.1:7860
	BaseURL string

	// API path is stored here
	Path *APIPath

	// If true, Field that starts with "Default" will be use everytime you make a request and if that field is empty/zero values
	UseDefault bool

	// Sampling Method or Sampler (e.g. Euler a, DPM++ 2M Karras). You can type it in yourself or use built-in Helper Package: sampler
	//
	// - Default: Euler a
	DefaultSampler string

	// Sampling Steps (e.g. 20, 120)
	//
	// - Default: 28
	DefaultSteps int

	// Classifier-Free Guidance Scale (e.g. 7, 12)
	//
	// - Default: 7
	DefaultCFGScale float64

	// Width of the image (e.g. 512, 1024)
	//
	// - Default: 512
	DefaultWidth int

	// Height of the image (e.g. 512, 1024)
	//
	// - Default: 512
	DefaultHeight int
}

type APIPath struct {
	// Path to txt2img API
	//
	// - Default: /sdapi/v1/txt2img
	Txt2Img string

	// Path to progress API
	//
	// - Default: /sdapi/v1/progress
	Progress string
}

var DefaultConfig = Config{
	BaseURL: "http://127.0.0.1:7860",
	Path: &APIPath{
		Txt2Img:  "/sdapi/v1/txt2img",
		Progress: "/sdapi/v1/progress",
	},
	DefaultSampler:  sampler.EULER_A,
	DefaultSteps:    28,
	DefaultCFGScale: 7,
	DefaultWidth:    512,
	DefaultHeight:   512,
}

func setDefault(conf ...Config) Config {
	if len(conf) <= 0 {
		return DefaultConfig
	}

	config := conf[0]

	if !config.UseDefault {
		return config
	}

	if config.BaseURL == "" {
		config.BaseURL = DefaultConfig.BaseURL
	}

	if config.Path == nil {
		config.Path = DefaultConfig.Path
	}

	if config.DefaultSampler == "" {
		config.DefaultSampler = DefaultConfig.DefaultSampler
	}

	if config.DefaultSteps == 0 {
		config.DefaultSteps = DefaultConfig.DefaultSteps
	}

	if config.DefaultCFGScale == 0 {
		config.DefaultCFGScale = DefaultConfig.DefaultCFGScale
	}

	if config.DefaultWidth == 0 {
		config.DefaultWidth = DefaultConfig.DefaultWidth
	}

	if config.DefaultHeight == 0 {
		config.DefaultHeight = DefaultConfig.DefaultHeight
	}

	return config
}
