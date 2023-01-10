package main

import (
	"log"

	"github.com/feo0o/kyanite/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
