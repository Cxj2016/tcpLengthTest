package main

import (
	"fmt"
	"os"
	"test/tcpTest/base"
	"time"
)

func main() {
	c := base.NewClient()
	if err := c.Connect(base.Addr); err != nil {
		fmt.Printf("client connect failed|err:%v\n", err)
		os.Exit(1)
	}

	for {
		msg := &base.Message{}
		for i := 0; i < 60000; i++ {
			msg.Data = append(msg.Data, 'a')
		}
		c.Write(msg)
		time.Sleep(time.Second * 10)
	}
}
