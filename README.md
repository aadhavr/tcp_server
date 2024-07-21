# Multi-threaded TCP Server in Go from Scratch

This is an educational project intended to help me understand how TCP servers work, using Go. If you're interested in learning how servers handle multiple client connections and process requests, you're in the right place.

## What's This All About?

Imagine a busy cafe. When you walk in, a barista takes your order, processes it, and eventually hands you your coffee. Now, picture this cafe as a server and each customer as a client. Our server (the cafe) can handle multiple clients (customers) at once, thanks to some clever coding.

## How It Works

### Conceptually

Our server opens a door, a TCP port, and waits for clients to knock (**Listening**). When a client knocks (connects), the server opens the door and lets them in (**Accepting**). The server reads the client's request, thinks about it (processes it), and sends back a response (**Processing**). Our server can handle multiple clients at the same time without making anyone wait too long. It's like having a barista be able to take multiple orders without serving one order's food, working in parallel. (**Concurrency**)

### In Code

The server listens on a specific port (like a doorbell waiting for someone to press it).It accepts incoming connections from clients (like opening the door for each customer). For each connection, it reads data, processes it (in our case, just waits a bit), and sends back a "Hello World" message. It uses goroutines (lightweight threads in Go) to handle multiple connections at the same time.

## Running the Server

**Get Go**

If you don't have Go installed, you'll need it. Download and install it from [golang.org](https://golang.org/dl/).

**Clone the Repo**

Download this project to your local machine.
```sh
git clone https://github.com/yourusername/simple-tcp-server.git
cd tcp_server
```

**Run the Server**

Start the server using Go.
```sh
go run main.go
```

**Test the Server**

Open another terminal and use `curl` to test it.
```sh
curl -v http://localhost:1729/
```
You should see a "Hello World" response.

**Load Testing**

To see how our server handles multiple requests, run the `load_test.sh` script included in the repo.
```sh
./load_test.sh
```

## Improvements and Enhancements

### Limiting the Number of Threads

To prevent the server from spawning too many goroutines, which can lead to high memory consumption and potentially crash the server. We use a semaphore or a buffered channel—a type of channel in Go that allows multiple values to be sent without a corresponding receiver being immediately ready to receive them, acting like a queue for data—to limit the number of concurrent threads. When a new connection is accepted, it acquires a slot from the semaphore. If no slots are available, the connection waits until a slot is freed. Once the goroutine finishes processing, it releases the slot back to the semaphore.

### Adding a Thread Pool

Creating and destroying threads (goroutines in Go) can be expensive in terms of CPU and memory. A thread pool reuses a fixed number of worker goroutines to handle incoming connections, reducing the overhead associated with thread management. We create a fixed number of worker goroutines at the start (the thread pool). These workers continuously listen for incoming connections from a job queue. When a connection is accepted, it's sent to the job queue, and an available worker processes it. This approach keeps the number of active goroutines constant and avoids the overhead of creating new ones for each connection.

### Connection Timeout

To avoid hanging connections that can block resources indefinitely. Without timeouts, a slow or unresponsive client can hold onto a connection, preventing the server from efficiently handling other requests. We set a deadline for read and write operations on each connection. If the operation doesn't complete within the specified time, it times out, and the connection is closed. This ensures that resources are freed up promptly, and the server can continue handling new connections.

## What You'll Learn

By exploring this project, you'll get a good grasp of:

- How TCP servers work conceptually.
- Basic Go programming for network applications.
- Handling multiple connections concurrently using goroutines.
- Practical experience with testing and load testing a server.

## Acknowledgements

Arpit Bhayani's [YouTube channel](https://www.youtube.com/@AsliEngineering) was really helpful for understanding these concepts.