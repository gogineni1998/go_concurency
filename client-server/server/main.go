package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handelconn(conn)
	}
}

func handelconn(conn net.Conn) {
	defer conn.Close()
	for {
		res, err := conn.Write([]byte("hi form sever\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}

}
