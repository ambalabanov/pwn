package pwn

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//Process func for connect
func Process(addr string) (c net.Conn) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Failed to connect: " + err.Error())
	}
	return
}

//Interactive IO
func Interactive(c net.Conn) {
	for {
		fmt.Printf("# ")
		sendData, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		c.Write([]byte(sendData))
		responseBuffer := make([]byte, 4096)
		numBytesRead, err := c.Read(responseBuffer)
		if err != nil {
			log.Printf("Error reading from server. " + err.Error())
			break
		}
		fmt.Printf("%s", responseBuffer[0:numBytesRead])
	}
}
