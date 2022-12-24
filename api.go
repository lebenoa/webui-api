package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type api struct {
	Config Config
}

func New(newConfig ...Config) *api {
	if len(newConfig) <= 0 {
		return &api{
			Config: Config{},
		}
	}

	return &api{
		Config: setDefault(newConfig[0]),
	}
}

// Convenience function to build prompt.
//
// BuildPrompt("masterpiece", "best quality", "solo") -> "masterpiece, best quality, solo"
func BuildPrompt(args ...string) string {
	return strings.Join(args, ", ")
}

func (a *api) get(path string) (body []byte, err error) {
	resp, err := http.Get(a.Config.BaseURL + path)
	if err != nil {
		return nil, err
	}

	return common(resp)
}

func (a *api) post(path string, data []byte) (body []byte, err error) {
	resp, err := http.Post(a.Config.BaseURL+path, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return common(resp)
}

func common(resp *http.Response) (body []byte, err error) {
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%v", string(body))
	}

	return
}
