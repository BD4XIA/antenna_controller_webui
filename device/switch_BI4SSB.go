package device

import (
	"fmt"
	"sync"
)

// BI4SSB antenna switch net controller
type Switch_BI4SSB struct {
	mu      sync.Mutex
	Address string
	conns   []int
}

func (sw *Switch_BI4SSB) Lock()   { sw.mu.Lock() }
func (sw *Switch_BI4SSB) Unlock() { sw.mu.Unlock() }

func (sw *Switch_BI4SSB) Load([]byte) {

}

func (sw *Switch_BI4SSB) Save() []byte {
	return nil
}

func (sw *Switch_BI4SSB) Info() string {
	return fmt.Sprintf("BI4SSB网络控制天线切换器(%s)", sw.Address)
}

func (sw *Switch_BI4SSB) GetConns() ([]int, error) {
	return sw.conns, nil
}

func (sw *Switch_BI4SSB) Connect(from int, to int) error {
	sw.conns[from] = to
	sw.conns[to] = from
	return nil
}

// func (sw *Switch) Query() ([][]byte, error) {
// 	data, err := sw.Get09("/JSON.txt")
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(data) < 2 {
// 		return nil, nil
// 	}

// 	reg, err := regexp.Compile(`\{.*?\}`)
// 	if err != nil {
// 		return nil, err
// 	}

// 	switches := reg.FindAll(data[1:len(data)-1], -1)
// 	return switches, nil
// }

// func (sw *Switch) EN(posw int) error {
// 	_, err := sw.Get09(fmt.Sprintf("/AT%d=ON", posw))
// 	return err
// }
