package main

import (
	"fmt"
	"net"
)

func handler(c net.Conn) {
	c.Write([]byte("hei David Saga"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	fmt.Println(l.Addr())
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(l.Addr())
			continue
		}
		go handler(c)
	}
}
