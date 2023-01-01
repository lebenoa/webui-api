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

	//  Auth contains username and password like this:
	//
	//   "username, password" or you can use Username && Password field instead.
	//
	//  But do keep in mind that if Auth is not empty string, it'll use Auth and not Username && Password
	Auth     string
	Username string
	Password string
}

var (
	httpClient = &http.Client{}
	API        *api
)

func New(newConfig ...Config) *api {
	API = &api{
		Config: setDefault(newConfig...),
	}
	return API
}

// Set username and password for use when making request. Equivalent to
//
//	api.Username = username
//	api.Password = password
//
// Either username or password should not be empty string
func (a *api) SetAuth(username, password string) {
	a.Username = username
	a.Password = password
}

// Convenience function to build prompt.
//
//	BuildPrompt("masterpiece", "best quality", "solo") -> "masterpiece, best quality, solo"
func BuildPrompt(args ...string) string {
	return strings.Join(args, ", ")
}

// Send Get Request.
func (a *api) get(path string) (body []byte, err error) {
	req, err := http.NewRequest("GET", a.Config.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}

	return a.exec(req)
}

// Send Post Request.
func (a *api) post(path string, data []byte) (body []byte, err error) {
	req, err := http.NewRequest("POST", a.Config.BaseURL+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return a.exec(req)
}

// Send Request.
func (a *api) exec(req *http.Request) ([]byte, error) {
	a.setAuth(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return readBody(resp)
}

// Set HTTP Basic Auth.
func (a *api) setAuth(req *http.Request) {
	if a.Auth != "" {
		credit := strings.Split(a.Auth, ", ")
		req.SetBasicAuth(credit[0], credit[1])
	} else if a.Username != "" && a.Password != "" {
		req.SetBasicAuth(a.Username, a.Password)
	}
}

// Read response body.
func readBody(resp *http.Response) (body []byte, err error) {
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
