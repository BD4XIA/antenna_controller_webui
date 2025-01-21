package device

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// BH4TDV's NET Rotate 2020 G1000
type Rotator_NEC2020_G1000 struct {
	mu      sync.Mutex
	Address string
	angle   int
}

func (rt *Rotator_NEC2020_G1000) Lock()   { rt.mu.Lock() }
func (rt *Rotator_NEC2020_G1000) Unlock() { rt.mu.Unlock() }

func (rt *Rotator_NEC2020_G1000) Load([]byte) {

}

func (rt *Rotator_NEC2020_G1000) Save() []byte {
	return nil
}

func (rt *Rotator_NEC2020_G1000) Info() string {
	return fmt.Sprintf("马工一轴网络控制旋转器(%s)", rt.Address)
}

func (rt *Rotator_NEC2020_G1000) GetElevation() (float32, error) { return rt.GetAzimuth() }
func (rt *Rotator_NEC2020_G1000) SetElevation(ele float32) error { return rt.SetAzimuth(ele) }
func (rt *Rotator_NEC2020_G1000) GetRoll() (float32, error)      { return rt.GetAzimuth() }
func (rt *Rotator_NEC2020_G1000) SetRoll(roll float32) error     { return rt.SetAzimuth(roll) }

func (rt *Rotator_NEC2020_G1000) GetAzimuth() (float32, error) {
	resp, err := http.Get("http://" + rt.Address + "/JSON.txt")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	log.Println(string(data))

	reg, err := regexp.Compile(`"G1000":"\s*\d{3}"`)
	if err != nil {
		return 0, err
	}

	angle, err := strconv.Atoi(strings.TrimLeft(strings.Split(string(reg.Find(data)), `"`)[3], " "))
	if err != nil {
		return 0, err
	}
	rt.angle = angle

	return float32(rt.angle), nil
}

func (rt *Rotator_NEC2020_G1000) SetAzimuth(az float32) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/index.html?G1000_F1=%03d", rt.Address, int(az)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reg, err := regexp.Compile(`\d{3} &deg`)
	if err != nil {
		return err
	}

	angle, err := strconv.Atoi(string(reg.Find(data)[:3]))
	if err != nil {
		return err
	}
	if angle != int(az) {
		return ErrSetFailed
	}

	r, err := http.Get("http://" + rt.Address + "/Q_K1=ON")
	if err != nil {
		return err
	}
	return r.Body.Close()
}

func (rt *Rotator_NEC2020_G1000) Update(az int, ele int, roll int) error {
	var cmd = ""

	if az > 0 {
		cmd = "/G1000=RIGHT"
	} else if az < 0 {
		cmd = "/G1000=LEFT"
	} else {
		cmd = "/G1000=STOP"
	}

	resp, err := http.Get("http://" + rt.Address + cmd)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
