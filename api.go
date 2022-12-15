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

// Covenient function to build prompt.
//
// BuildPrompt("masterpiece", "best quality", "solo", "cute", "innocent") -> "masterpiece, best quality, solo, cute, innocent"
func BuildPrompt(args ...string) string {
	return strings.Join(args, ", ")
}

func (a *api) post(path string, data []byte) ([]byte, error) {
	resp, err := http.Post(a.Config.BaseURL+path, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
