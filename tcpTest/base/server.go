package base

import (
	"fmt"
	"net"
)

type TcpServer struct {
	listener net.Listener
}

func NewTcpServer() *TcpServer {
	return &TcpServer{}
}

func (t *TcpServer) Listen(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("listen failed|err:%v\n", err)
		return err
	}

	t.listener = lis
	go t.acceptConn()

	return err
}

func (t *TcpServer) acceptConn() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("accept failed|err:%v\n", err)
			return
		}

		go t.handleConn(conn)
	}
}

func (t *TcpServer) handleConn(c net.Conn) {
	newConn := NewConnection(c)
	newConn.Run()
}
