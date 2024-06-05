package util

import (
	"bytes"
	"fmt"
	"net"
)

func Start_server(port string) (net.Conn, error) {

	conn, err := net.Dial("tcp", port)

	if err != nil {

		return nil, fmt.Errorf("error connecting to server not able to connect")

	}

	return conn, nil

}

func Send_data(data []byte, connection net.Conn) error {

	_, err := connection.Write(data)

	if err != nil {
		return fmt.Errorf("error sendind data: %v", err)
	}

	return nil
}

func Recive_data(connect net.Conn) ([]byte, error) {

	buffer := make([]byte, 1024)
	var result bytes.Buffer

	for {
		n, err := connect.Read(buffer)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}

		result.Write(buffer[:n])

		if bytes.Contains(buffer[:n], []byte{'\n'}) {
			break
		}
	}

	return result.Bytes(), nil

}
