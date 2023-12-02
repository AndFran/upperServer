package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		clientReader := bufio.NewReader(conn) // read from client
		clientText, err := clientReader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		_, err = fmt.Fprintln(conn, strings.ToUpper(clientText)+"\n")
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	fmt.Println("Welcome to UPPER server....waiting for a connection")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("Accepted connection")
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
