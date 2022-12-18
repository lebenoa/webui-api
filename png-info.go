package api

import (
	"fmt"

	"github.com/goccy/go-json"
)

func (a *api) PNGInfo(image string) (string, error) {
	payload, err := json.Marshal(map[string]string{"image": image})
	if err != nil {
		return "", err
	}
	fmt.Println(string(payload))

	data, err := a.post(a.Config.Path.PNGInfo, payload)
	if err != nil {
		return "", err
	} else if string(data) == "Internal Server Error" {
		return "", fmt.Errorf("%v", "500 internal server error")
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", fmt.Errorf("unmarshaling json: %v\nValue: %v", err, string(data))
	}

	info, ok := result["info"].(string)
	if !ok {
		return "", fmt.Errorf("%v", string(data))
	}

	return info, err
}
