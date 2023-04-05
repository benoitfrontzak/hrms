package main

import "log"

// Customize prefix log
func setLog() {
	log.SetPrefix("[Front-end]: ")
	log.SetFlags(0) // remove file:line and timestamps from log liness
}
