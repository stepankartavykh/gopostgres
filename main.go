package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mypostgres/query"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func postmasterProcess() {
	// TODO How postmaster process is launching in PostgresSQL postmaster?
	for {
		// Simulation of work related to processes.
		// TODO Figure out what processes are working and what logic behind them (e.g. how connections are handled).
		time.Sleep(time.Second)
		fmt.Printf("Postmaster process (%d) is working!\n", os.Getpid())
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("DATABASE_PORT")
	dirWithDatabasesPath := os.Getenv("DATABASE_LOCATION_DIRECTORY_PATH")
	createInitialDatabase(dirWithDatabasesPath)

	fmt.Println(os.Getenv("GOPATH"))

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("listening on %s\n", listener.Addr())
	go postmasterProcess()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func createInitialDatabase(dirPath string) {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		fmt.Printf("Folder %s for initial database does not exist.\n", dirPath)
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
		// TODO Figure out how to create logging in GO.
		response := query.HandleQuery(string(requestQuery))
		line := response.GetResponseResult()
		write, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(write)
		}
	}
}
