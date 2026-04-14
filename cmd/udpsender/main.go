package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("error", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("err reading:", err)
			continue
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Println("error writing: ", err)
			return
		}

	}
}
