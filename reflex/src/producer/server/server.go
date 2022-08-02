package server

import (
	"context"
	"log"

	"github.com/neurotempest/mq_test/reflex/src/producer/pb"
	"github.com/neurotempest/mq_test/reflex/src/producer/state"
)

var _ pb.ProducerServer = (*server)(nil)

type server struct {
	pb.UnimplementedProducerServer
	state state.State
}

func New(st state.State) server {

	return server{
		state: st,
	}
}

func (s *server) Ping(
	ctx context.Context,
	req *pb.PingRequest,
) (*pb.PingResponse, error) {

	log.Println("Producer.Ping!!")

	return &pb.PingResponse{}, nil
}
