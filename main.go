package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	defer conn.Close() // 8. ensure the connection is closed when done

	// 9. set a deadline for read/write operations to prevent hanging connections
	conn.SetDeadline(time.Now().Add(10 * time.Second))

	buf := make([]byte, 1024) // 10. create a buffer to read data into
	_, err := conn.Read(buf)  // 11. read data from the connection into the buffer
	if err != nil {
		log.Printf("read error: %v\n", err) // 12. log any read error
		return
	}

	log.Println("processing the request")
	time.Sleep(1 * time.Second) // 13. simulate processing time

	// 14. write the response back to the client
	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nhello world\r\n"))
	if err != nil {
		log.Printf("write error: %v\n", err) // 15. log any write error
		return
	}
}

func worker(jobQueue <-chan net.Conn) {
	for conn := range jobQueue { // 7. continuously receive connections from the job queue
		do(conn) // 8. handle each connection
	}
}

func main() {
	// 1. create a tcp listener on port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err) // 2. log and exit if there's an error creating the listener
	}
	defer listener.Close() // 17. ensure the listener is closed on exit

	log.Println("server started on :1729")

	// 3. define the pool size and create a job queue channel
	const poolSize = 10
	jobQueue := make(chan net.Conn, poolSize)

	// 4. start worker goroutines to handle connections from the job queue
	for i := 0; i < poolSize; i++ {
		go worker(jobQueue) // 5. launch a worker goroutine
	}

	// 6. define the maximum number of concurrent threads and create a semaphore channel
	const maxThreads = 10
	sem := make(chan struct{}, maxThreads)

	for {
		log.Println("waiting for a client to connect")
		conn, err := listener.Accept() // 16. accept a new client connection
		if err != nil {
			log.Printf("accept error: %v\n", err) // 17. log any accept error
			continue
		}

		sem <- struct{}{} // 18. acquire a slot in the semaphore
		go func(conn net.Conn) {
			defer func() { <-sem }() // 21. release the slot in the semaphore when done
			jobQueue <- conn         // 19. send the connection to the job queue
		}(conn) // 20. starts a new goroutine to handle the connection
	}
}
