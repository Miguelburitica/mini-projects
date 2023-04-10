package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string // an outgoing message channel

var (
	entering = make(chan Client)
	leaveing = make(chan Client)
	chatMessages = make(chan string) // all incoming client messages
)

var (
	host = flag.String("h", "localhost", "host to scan")
	port = flag.Int("p", 3000, "port to scan")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	clientMessages := make(chan string) // outgoing client messages
	go MessageWriter(conn, clientMessages)

	who := conn.RemoteAddr().String()
	clientMessages <- "You are " + who
	chatMessages <- who + " has arrived"
	entering <- clientMessages

	input := bufio.NewScanner(conn)
	for input.Scan() {
		chatMessages <- who + ": " + input.Text()
	}

	leaveing <- clientMessages
	chatMessages <- who + " has left"
}

func MessageWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func Broadcast() {
	clients := make(map[Client]bool) // all connected clients
	for {
		select {
		case msg := <-chatMessages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for client := range clients {
				client <- msg
			}

		case client := <-entering:
			clients[client] = true

		case client := <-leaveing:
			delete(clients, client)
			close(client)
		}
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))

	if err != nil {
		log.Fatal(err)
	}
	go Broadcast() // start the broadcaster

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn) // handle connections concurrently
	}
}