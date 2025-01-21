package device

import (
	"fmt"
	"regexp"
)

type Switch struct {
	address string
}

func NewSwitch(ip string, port int) *Switch {
	address := fmt.Sprintf("%s:%d", ip, port)
	return &Switch{address}
}

func (sw *Switch) Query() ([][]byte, error) {
	data, err := sw.Get09("/JSON.txt")
	if err != nil {
		return nil, err
	}
	if len(data) < 2 {
		return nil, nil
	}

	reg, err := regexp.Compile(`\{.*?\}`)
	if err != nil {
		return nil, err
	}

	switches := reg.FindAll(data[1:len(data)-1], -1)
	return switches, nil
}

func (sw *Switch) EN(port int) error {
	_, err := sw.Get09(fmt.Sprintf("/AT%d=ON", port))
	return err
}
