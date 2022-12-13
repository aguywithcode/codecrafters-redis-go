package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// create the server
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	//Handle connections in an event loop.
	for {
		conn, err := l.Accept()
		go handleConnection(err, conn)
	}
}

func handleConnection(err error, conn net.Conn) {
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connection accepted...waiting for ping")
	w := bufio.NewWriter(conn)
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		switch err {
		case nil:
			fmt.Println("Ping received, sending response", string(buf))
			break
		case io.EOF:
			fmt.Println("Eof received")
			return
		default:
			fmt.Println("ERROR", err)
		}
		if _, err := w.WriteString("+PONG\r\n"); err != nil {
			fmt.Println("Error writing response: ", err.Error())
		}
		fmt.Println("Flushing the buffer")
		w.Flush()
	}
}
