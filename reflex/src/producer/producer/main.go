package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.Println("Hello world from producer")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-ch:
		log.Println("Received OS signal:", sig.String())
		return
	}
}

