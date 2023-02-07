package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/feo0o/kyanite/server"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("caught SIGTERM, shutting down")
		os.Exit(0)
	}()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
