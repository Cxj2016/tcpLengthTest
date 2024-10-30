package base

import (
	"fmt"
	"net"
)

type TcpClient struct {
	conn *Connection
}

func NewClient() *TcpClient {
	return &TcpClient{}
}

func (t *TcpClient) Connect(address string) error {
	cn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("client dial failed|err:%v\n", err)
		return err
	}
	t.conn = NewConnection(cn)
	t.conn.Run()

	return nil
}

func (t *TcpClient) Write(msg *Message) {
	t.conn.writeQ <- msg
}
