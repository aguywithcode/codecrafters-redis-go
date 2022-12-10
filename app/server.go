package main

import (
	"bufio"
	"fmt"
	"io"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connection accepted...waiting for ping")
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	for {
		_, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Println("Ping received, sending response")
			break
		case io.EOF:
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
