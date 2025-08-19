package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// Get target host and port from environment variables or use defaults
	targetHost := os.Getenv("PROXY_TARGET_HOST")
	if targetHost == "" {
		targetHost = "mc.hypixel.net"
	}
	
	targetPort := os.Getenv("PROXY_TARGET_PORT")
	if targetPort == "" {
		targetPort = "25565"
	}
	
	listenAddr := ":25565"
	
	log.Printf("Starting Minecraft proxy to %s:%s", targetHost, targetPort)
	
	// Start listening for connections
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()
	
	log.Printf("Proxy listening on %s", listenAddr)
	
	// Accept connections
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		
		// Handle each connection in a goroutine
		go handleConnection(clientConn, targetHost, targetPort)
	}
}

func handleConnection(clientConn net.Conn, targetHost, targetPort string) {
	defer clientConn.Close()
	
	// Connect to the target server
	serverConn, err := net.Dial("tcp", targetHost+":"+targetPort)
	if err != nil {
		log.Printf("Failed to connect to target server: %v", err)
		return
	}
	defer serverConn.Close()
	
	log.Printf("Proxying connection from %s to %s:%s", 
		clientConn.RemoteAddr(), targetHost, targetPort)
	
	// Copy data between client and server
	go io.Copy(serverConn, clientConn)
	io.Copy(clientConn, serverConn)
	
	log.Printf("Connection from %s closed", clientConn.RemoteAddr())
}