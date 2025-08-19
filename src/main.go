package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

// ready indicates whether the service is ready to serve traffic
var ready int32

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
	
	// Start the health check server in a separate goroutine
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if atomic.LoadInt32(&ready) == 1 {
				fmt.Fprintf(w, "Minecraft Proxy is running and ready")
			} else {
				http.Error(w, "Service not ready", http.StatusServiceUnavailable)
			}
		})
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			if atomic.LoadInt32(&ready) == 1 {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "OK")
			} else {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, "Not ready")
			}
		})
		log.Println("Health check server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
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
	
	// Mark the service as ready
	atomic.StoreInt32(&ready, 1)
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