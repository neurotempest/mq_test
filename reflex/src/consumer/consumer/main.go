package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
)

func main() {

	flag.Parse()

	log.Println("Hello world from consumer")

	st := state.New()

	log.Println("Created consumer state...")

	ctx := context.TODO()

	err := st.GetProducerClient().Ping(ctx, "pinger from consumer")
	if err != nil {
		log.Fatal("ping err:", err.Error())
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-ch:
		log.Println("Received OS signal:", sig.String())
		return
	}
}


