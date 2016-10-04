package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	/* Lets prepare a address at any address at port 10001*/
	serverIP := flag.String("serverIP", "127.0.0.1", "Server IP")
	serverPort := flag.String("serverPort", "20001", "Server Port")
	flag.Parse()
	serverAddr := *serverIP + ":" + *serverPort
	ServerAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Time:", time.Now(), "Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
