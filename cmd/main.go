package main

import (
	"log"

	"livestream-app/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
