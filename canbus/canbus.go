package canbus

import (
	"io"
	"net"

	"golang.org/x/sys/unix"
)

type device struct {
	fd int
}

func (d device) Read(data []byte) (int, error) {
	return unix.Read(d.fd, data)
}

func (d device) Write(data []byte) (int, error) {
	return unix.Write(d.fd, data)
}

type Socket struct {
	iface *net.Interface
	addr  *unix.SockaddrCAN
	dev   device
}

func New() (*Socket, error) {
	fd, err := unix.Socket(unix.AF_CAN, unix.SOCK_RAW, unix.CAN_RAW)
	if err != nil {
		return nil, err
	}
	return &Socket{dev: device{fd}}, nil
}

func (sck *Socket) Close() error {
	return unix.Close(sck.dev.fd)
}

func (sck *Socket) Bind(addr string) error {
	iface, err := net.InterfaceByName(addr)
	if err != nil {
		return err
	}
	sck.iface = iface
	sck.addr = &unix.SockaddrCAN{Ifindex: sck.iface.Index}

	return unix.Bind(sck.dev.fd, sck.addr)
}

func (sck *Socket) RecvRawFrame() (msg Frame, err error) {

	var frame [frameSize]byte
	n, err := io.ReadFull(sck.dev, frame[:])
}
