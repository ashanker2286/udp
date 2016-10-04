package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	serverIP := flag.String("serverIP", "127.0.0.1", "Server IP")
	serverPort := flag.String("serverPort", "20001", "Server Port")
	clientIP := flag.String("clientIP", "127.0.0.1", "Client IP")
	clientPort := flag.String("clientPort", "20002", "Client Port")
	flag.Parse()
	serverAddr := *serverIP + ":" + *serverPort
	ServerAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	CheckError(err)

	clientAddr := *clientIP + ":" + *clientPort
	LocalAddr, err := net.ResolveUDPAddr("udp", clientAddr)
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		fmt.Println("Time: ", time.Now(), "Sending:", msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Millisecond * 250)
	}
}
