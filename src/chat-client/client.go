package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))
		bs := make([]byte, 1024)
		n, err := conn.Read(bs)
		if err != nil {
			panic(err)
		}
		fmt.Printf("received : %s", string(bs[:n]))
	}

}
