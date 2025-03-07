package main

import (
	"queueAppV2/internal/app/local"

)

func main() {
	server := local.New()
	if err := server.Run(); err != nil {
		panic(err)
	}
}