package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	dialer, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer dialer.Close()

	for {
		clientInputReader := bufio.NewReader(os.Stdin)
		fmt.Print(">>")
		text, _ := clientInputReader.ReadString('\n')
		_, err = fmt.Fprintln(dialer, text+"\n") // writer to server
		if err != nil {
			log.Fatal(err)
		}

		serverReader := bufio.NewReader(dialer)
		serverResponse, _ := serverReader.ReadString('\n')
		fmt.Println(serverResponse)
	}
}
