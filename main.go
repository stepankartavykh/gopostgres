package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Args[1])
	prefix := os.Args[2]
	fmt.Println(os.Args)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("listening on %s, prefix: %s\n", listener.Addr(), prefix)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		go handleConnection(conn, prefix)
	}
}

func handleConnection(conn net.Conn, prefix string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(conn.LocalAddr())
			fmt.Println(conn.RemoteAddr())
		}
	}(conn)
	reader := bufio.NewReader(conn)
	for {
		requestQuery, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", requestQuery)
		handleQuery(string(requestQuery))
		line := fmt.Sprintf("%s %s", prefix, requestQuery)
		fmt.Printf("response: %s", line)
		write, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(write)
		}
	}
}
