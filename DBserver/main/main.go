package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	fmt.Println("Listening on port 8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		return
	}
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		conn.Write([]byte("+OK\r\n"))
	}
}
