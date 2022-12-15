package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"

	"github.com/Meonako/webui-api/utils"
)

type progressRespond struct {
	Progress     float64 `json:"progress"`
	ETA          float64 `json:"eta_relative"`
	State        state   `json:"state"`
	CurrentImage string  `json:"current_image"`
}

type state struct {
	Skipped             bool   `json:"skipped"`
	Interrupted         bool   `json:"interrupted"`
	Job                 string `json:"job"`
	JobCount            int    `json:"job_count"`
	JobNo               int    `json:"job_no"`
	CurrentStep         int    `json:"sampling_step"`
	TargetSamplingSteps int    `json:"sampling_steps"`
}

// Get Generation Progress Info. Return Respond struct and Error object.
func (a *api) Progress() (*progressRespond, error) {
	resp, err := a.get(a.Config.Path.Progress)
	if err != nil {
		return &progressRespond{}, err
	}

	var result progressRespond
	err = json.Unmarshal(resp, &result)
	return &result, err
}

// Convert Progress field to percentage value (e.g. 83.57%).
//
// format argument is optional. pass emtpy string will round to interger in string format.
//
// Default: 2 (e.g. 83.57%)
func (p *progressRespond) GetProgress(format ...string) string {
	deci := ".2"
	if len(format) > 0 {
		if format[0] == "" {
			deci = ".0"
		} else {
			deci = "." + format[0]
		}
	}

	return fmt.Sprintf("%"+deci+"f", p.Progress*100) + "%"
}

// "Estimated Time of Arrival" of images :) (e.g. 53 Minutes 10 Seconds)
//
//	format argument is optional. pass "2" will round "Seconds" to 2 decimal places in string format.
//
//	Default: 0 (e.g. 53 Minutes 10 Seconds)
func (p *progressRespond) GetETA(format ...string) string {
	deci := ".0"
	if len(format) > 0 {
		if format[0] != "" {
			deci = "." + format[0]
		}
	}

	return fmt.Sprintf("%v Minutes %"+deci+"f Seconds", math.Floor(p.ETA/60), math.Mod(p.ETA, 60))
}

// Get current image if not empty. Usually this is empty if you didn't change the value of the following setting.
//
// "Show image creation progress every N sampling steps. Set to 0 to disable. Set to -1 to show after completion of batch."
func (p *progressRespond) GetCurrentImage() (*bytes.Reader, error) {
	if p.CurrentImage == "" {
		return nil, nil
	}

	byte, err := utils.DecodeBase64(p.CurrentImage)

	return bytes.NewReader(byte), err
}
