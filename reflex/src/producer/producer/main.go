package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/neurotempest/mq_test/reflex/src/producer/pb"
	"github.com/neurotempest/mq_test/reflex/src/producer/server"
	"github.com/neurotempest/mq_test/reflex/src/producer/state"
)

var grpcAddr = flag.String("grpc_address", ":1234", "host:port to server gRPC service")

func main() {

	flag.Parse()

	log.Println("Hello world from producer")

	st := state.New()

	grpcStopFn := ServeGRPCForever(st)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-ch:
		log.Println("Received OS signal:", sig.String())
		grpcStopFn()
		return
	}
}

func ServeGRPCForever(st state.State) func() {
	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := server.New(st)

	pb.RegisterProducerServer(grpcServer, &srv)
	go func() {
		err := grpcServer.Serve(lis)
		log.Fatal("grpc server exited:", err.Error())
	}()

	return func() {
		log.Println("waiting for grpc server to stop...")
		grpcServer.GracefulStop()
	}
}
