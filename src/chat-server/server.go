package main

import (
	"net"
	"fmt"
	"bufio"
)

func acceptUsers(ln net.Listener, user chan net.Conn, dcons chan net.Conn) {
	for {
		conn, err := ln.Accept()
		fmt.Printf("Received Connection\n")
		if err != nil {
			panic(err)
		}
		user <- conn
	}
}

func readMessages(accons map[net.Conn]int, conn net.Conn, dcons chan net.Conn, messages chan string) {
	rd := bufio.NewReader(conn)
	for {
		m, err := rd.ReadString('\n')
		if err != nil {
			panic(err)
		}
		messages <- fmt.Sprintf("Client %v : %v", accons[conn], m)
	}

	dcons <- conn
}

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}
	conn := make(chan net.Conn, 0)
	dcons := make(chan net.Conn, 0)
	message := make(chan string)
	accons := make(map[net.Conn]int)
	i := 0
	go acceptUsers(ln, conn, dcons)

	for {
		select {
		case u := <-conn:
			accons[u] = i
			i++
			go readMessages(accons, u, dcons, message)
		case msg := <-message:
			for c := range accons {
				c.Write([]byte(msg))
			}
		case d := <-dcons:
			fmt.Printf("Client %v disconnected", d)
			delete(accons, d)
		}
	}
}
