package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
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
	
	// Determine the port for the health check server
	// Render typically sets the PORT environment variable for web services
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default to port 80 if PORT is not set
	}
	
	// Start the health check server in a separate goroutine
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Minecraft Proxy is running")
		})
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "OK")
		})
		log.Printf("Health check server starting on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Printf("Health server error: %v", err)
		}
	}()
	
	// Give the health server a moment to start
	time.Sleep(100 * time.Millisecond)
	
	// Start the proxy
	listenAddr := ":25565"
	log.Printf("Starting Minecraft proxy to %s:%s", targetHost, targetPort)
	
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()
	
	log.Printf("Proxy listening on %s - service is ready", listenAddr)
	
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