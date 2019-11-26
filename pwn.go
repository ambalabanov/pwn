package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const addr = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Failed to connect: " + err.Error())
	}
	defer conn.Close()
	for {
		fmt.Printf("# ")
		sendData, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(sendData))
		responseBuffer := make([]byte, 4096)
		numBytesRead, err := conn.Read(responseBuffer)
		if err != nil {
			log.Printf("Error reading from server. " + err.Error())
			break
		}
		fmt.Printf("%s", responseBuffer[0:numBytesRead])
	}
}
