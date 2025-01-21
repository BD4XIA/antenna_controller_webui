package device

import (
	"io"
	"net"
)

func (sw *Switch) Get09(path string) ([]byte, error) {
	conn, err := net.Dial("tcp", sw.address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	request := "GET " + path
	_, err = conn.Write([]byte(request))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(conn)
}
