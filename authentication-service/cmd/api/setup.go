package main

import "log"

// Customize prefix log
func setLog() {
	log.SetPrefix("[AUTH-Service]: ")
	log.SetFlags(0) // remove file:line and timestamps from log liness
}
