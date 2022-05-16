package ziface

import "net"

type IStream interface {
	Close() error
	RemoteAddr() net.Addr
	Write(b []byte) (int, error)
	Read(p []byte) (n int, err error)
}
