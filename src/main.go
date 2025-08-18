package main

import (
	"log"
	"net"

	"github.com/haveachin/go-mcproxy"
)

func main() {
	// Configuration for the proxy
	proxy := mcproxy.NewProxy()
	proxy.ListenAddr = ":25565"
	proxy.TargetAddr = "mc.hypixel.net:25565"

	// Start the proxy
	log.Println("Starting Minecraft proxy to mc.hypixel.net:25565")
	if err := proxy.Start(); err != nil {
		log.Fatalf("Failed to start proxy: %v", err)
	}
}