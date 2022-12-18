package api

import (
	"bytes"
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

func (a *api) get(path string) ([]byte, error) {
	resp, err := http.Get(a.Config.BaseURL + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (a *api) post(path string, data []byte) ([]byte, error) {
	resp, err := http.Post(a.Config.BaseURL+path, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
