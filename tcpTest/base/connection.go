package base

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

type Connection struct {
	conn   net.Conn
	writeQ chan *Message
	writer *bufio.Writer
	reader *bufio.Reader
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn:   conn,
		writeQ: make(chan *Message, 100),
		writer: bufio.NewWriterSize(conn, 1024*64),
		reader: bufio.NewReaderSize(conn, 1024*64),
	}
}

func (c *Connection) Run() {
	go c.loopRead()
	go c.loopWrite()
}

func (c *Connection) loopRead() {
	for {
		header, err := c.reader.Peek(HeaderLength)
		if err != nil {
			fmt.Printf("peek header failed|err:%v\n", err)
			return
		}

		dataLength := binary.BigEndian.Uint32(header[0:])
		fmt.Printf("read data to read|header:%v|length:%d\n", header, dataLength)

		msgBuff := make([]byte, HeaderLength+dataLength)
		if _, err = c.reader.Read(msgBuff); err != nil {
			fmt.Printf("read msg buff failed|err:%v\n", err)
			return
		}

		fmt.Printf("read total length:%d|data:%v\n", len(msgBuff), msgBuff[HeaderLength:])
	}
}

func (c *Connection) loopWrite() {
	for {
		select {
		case msg := <-c.writeQ:
			var header = make([]byte, HeaderLength) //数据长度
			length := len(msg.Data)
			binary.BigEndian.PutUint32(header[0:], uint32(length))
			fmt.Printf("write header:%v|length:%d\n", header, length)
			if _, err := c.writer.Write(header); err != nil {
				fmt.Printf("write header failed|err:%v\n", err)
				return
			}

			if length > 0 {
				if _, err := c.writer.Write(msg.Data); err != nil {
					fmt.Printf("write body failed|err:%v\n", err)
					return
				}
			}

			if err := c.writer.Flush(); err != nil {
				fmt.Printf("flush failed|err:%v\n", err)
				return
			}
		}
	}
}
