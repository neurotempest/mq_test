package server

import (
	"context"

	"github.com/neurotempest/mq_test/reflex/src/consumer/pb"
	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
)

var _ pb.ConsumerServer = (*server)(nil)

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

	return nil, nil
}
