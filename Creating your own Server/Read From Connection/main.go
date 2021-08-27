package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	Scanner := bufio.NewScanner(conn)
	for Scanner.Scan() {
		ln := Scanner.Text()
		fmt.Println(ln)
	}
	//This program will keep on running as the program gets stuck on the upper loop
	defer conn.Close()
	fmt.Println("Code got here!")
}
