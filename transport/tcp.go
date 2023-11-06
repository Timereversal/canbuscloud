package transport

import (
	"github.com/go-daq/canbus"
	"net"
)

type Transport struct {
	address    string
	connection net.Conn
}

func New(addr string) (*Transport, error) {
	var t *Transport
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return t, err
	}
	t.connection = conn
	t.address = addr
	return t, nil
}

func (t *Transport) send(frames []canbus.Frame) {
	//several canbus.Frame will be delivered
	//	need to create a protocol to send more than 1 canbus.Frame
}
