package main

import (
	"fmt"
	"log"
	"net"
)

func clientSend() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to server.")
	var message string

	for{
		fmt.Print("Enter message to send: ")
		fmt.Scanf("%s", &message)
		if message == "exit" {
			break
		}
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatalf("Error sending data: %v", err)
		}
		fmt.Printf("Sent: %s\n", message)

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Error reading response: %v", err)
		}
		fmt.Printf("Received response: %s\n", string(buffer[:n]))
	}
}