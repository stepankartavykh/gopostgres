package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"mypostgres/query"
	"net"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("DATABASE_PORT")

	fmt.Println(os.Getenv("GOPATH"))

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("listening on %s\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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
		query.HandleQuery(string(requestQuery))
		line := fmt.Sprintf("%s", requestQuery)
		fmt.Printf("response: %s", line)
		write, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(write)
		}
	}
}
