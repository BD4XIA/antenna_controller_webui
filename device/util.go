package device

import (
	"io"
	"net"
)

func get09(address string, path string) ([]byte, error) {
	conn, err := net.Dial("tcp", address)
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
