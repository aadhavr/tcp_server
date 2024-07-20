package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) { // the connection is read, responded, and closed
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("processing the request")
	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for { // this loop keeps the server open, despite multiple requests
		log.Println("waiting for a client to connect")
		conn, err := listener.Accept() // the tcp server waits for a client to connect
		if err != nil {
			log.Fatal(err)
		}

		go do(conn) // once the client connects, it does a connection
		// the go before "do" ensures that the moment the request is sent it accepts it and starts again.
		// remove it and it becomes a singlethreaded server
	}
}
