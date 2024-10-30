package main

import (
	"fmt"
	"os"
	"test/tcpTest/base"
)

func main() {
	server := base.NewTcpServer()
	if err := server.Listen(base.Addr); err != nil {
		fmt.Printf("listen failed|err:%v\n", err)
		os.Exit(1)
	}
	for {

	}
}
