package api

import (
	"github.com/goccy/go-json"
)

type stableDiffusionModels struct {
	Title     string `json:"title"`
	ModelName string `json:"model_name"`
	Hash      string `json:"hash"`
	Filename  string `json:"filename"`
	Config    string `json:"config"`
}

// Get available Stable Diffusion Models
func (a *api) SDModels() (result []*stableDiffusionModels, err error) {
	resp, erro := a.get(a.Config.Path.SDModels)
	if erro != nil {
		err = erro
		return
	}

	err = json.Unmarshal(resp, &result)
	return
}
