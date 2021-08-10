package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

var clientCounter uint

func main() {
	var clients []net.Conn

	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:5555")
	if err != nil {
		log.Fatalln(err)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpListener.Close()

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		clients = append(clients, conn)
		clientCounter++

		fmt.Printf("%v connected\n", conn.RemoteAddr().String())

		go handleClient(conn)
	}
}

func handleClient(c net.Conn) {
	_, err := c.Write([]byte("hello client " + strconv.Itoa(int(clientCounter))))
	if err != nil {
		log.Fatalln("couldnt send message to client with ip: ", c.RemoteAddr().String())
	}

	// send byte array of kittens.mp4 to the client
	kittensVideo, err := os.ReadFile("kittens.mp4")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("sending file...")
	_, _ = c.Write(kittensVideo)
	fmt.Println("sent.")
}