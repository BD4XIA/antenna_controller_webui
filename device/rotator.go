package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Rotator struct {
	address string
}

type RotatorStatus int

func NewRotator(ip string, port int) *Rotator {
	address := fmt.Sprintf("%s:%d", ip, port)
	return &Rotator{address}
}

type jsonRotator struct {
	Setup []JsonSetup `json:"setup"`
}

type JsonSetup struct {
	G1000 string `json:"G1000,omitempty"`
	Stu   string `json:"STU,omitempty"`
}

func (rt *Rotator) Status() ([]JsonSetup, error) {
	resp, err := http.Get("http://" + rt.address + "/JSON.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var status jsonRotator
	err = json.Unmarshal(data, &status)
	if err != nil {
		return nil, err
	}

	return status.Setup, nil
}

func (rt *Rotator) Query() ([]string, error) {
	resp, err := http.Get("http://" + rt.address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reg, err := regexp.Compile(`[0-9]{3} &deg`)
	if err != nil {
		return nil, err
	}

	angles := reg.FindAll(data, -1)
	if len(angles) != 6 {
		return nil, nil
	}

	var results [6]string
	for i, angle := range angles {
		results[i] = string(angle[:3])
	}

	return results[:], nil
}

func (rt *Rotator) Set(angles []string) error {
	var config [6]string
	for i, angle := range angles {
		config[i] = fmt.Sprintf("G1000_F%d=%s", i+1, angle)
	}

	resp, err := http.Get("http://" + rt.address + "/index.html?" + strings.Join(config[:], "&"))
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func (rt *Rotator) EN(id int) error {
	resp, err := http.Get("http://" + rt.address + fmt.Sprintf("/Q_K%d=ON", id))
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
