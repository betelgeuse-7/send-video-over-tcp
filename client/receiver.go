// send video to the server over a tcp connection
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp", "localhost:5555")
	if err != nil {
		log.Fatalln(err)
	}
	
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatalln(err)
	}

	video := make([]byte, 2500000)

	for {
		n, _ := conn.Read(video)
		if n >= len(video) {
			break
		}
	}

	fmt.Println("got: ", video)

	// create the file and write the bytes to it
	f, err := os.Create(generateRandomFileName()+".mp4")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = f.Write(video)
	if err != nil {
		log.Fatalln(err)
	}
}

func generateRandomFileName() string {
	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
