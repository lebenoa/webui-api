package api

import "github.com/goccy/go-json"

type Interrogate struct {
	Image string `json:"image"`
	Model string `json:"model"`
}

// Get captions from an image.
func (a *api) Interrogate(params *Interrogate) (string, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	resp, err := a.post(a.Config.Path.Interrogate, payload)
	if err != nil {
		return "", err
	}

	result := map[string]string{}
	json.Unmarshal(resp, &result)

	return result["caption"], nil
}
