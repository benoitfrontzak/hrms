package main

import (
	"broker/pkg/web"
	"log"
)

const webPort = "8088"

func main() {
	// Start web server
	if err := web.Run(multiplexer(), webPort); err != nil {
		log.Fatalf(err.Error())
	}
}
