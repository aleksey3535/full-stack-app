package main

import "queueAppV2/internal/app/dev"

func main() {
	server := dev.New()
	if err := server.Run(); err != nil {
		server.Stop()
	}
}