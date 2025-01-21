package device

import (
	"errors"
	"log"
)

type AtomicDevice interface {
	Lock()
	Unlock()
	Load([]byte)
	Save() []byte
	Info() string
}

type Rotator interface {
	AtomicDevice
	GetAzimuth() (float32, error)
	SetAzimuth(float32) error
	GetElevation() (float32, error)
	SetElevation(float32) error
	GetRoll() (float32, error)
	SetRoll(float32) error
	Update(int, int, int) error
}

func InitRotator(rt Rotator) (Rotator, error) {
	switch rt.(type) {
	case *Rotator_NEC2020_G1000:
		log.Printf("initing rotator: %s", rt.Info())
	default:
		return nil, ErrUnknownDevice
	}
	return rt, nil
}

type Switch interface {
	AtomicDevice
	GetConns() ([]int, error)
	Connect(int, int) error
}

func InitSwitch(sw Switch) (Switch, error) {
	switch sw.(type) {
	case *Switch_BI4SSB:
		log.Printf("initing switch: %s", sw.Info())
	default:
		return nil, ErrUnknownDevice
	}
	return sw, nil
}

var ErrUnknownDevice = errors.New("unknown device")
var ErrIllegalData = errors.New("illegal data")
var ErrSetFailed = errors.New("failed to set device")
var ErrExhaustedConn = errors.New("connections exhausted")
