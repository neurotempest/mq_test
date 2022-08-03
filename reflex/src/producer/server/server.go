package server

import (
	"context"

	"github.com/neurotempest/mq_test/reflex/src/producer/ops"
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

	err := ops.Ping(ctx, s.state, req.Msg)
	if err != nil {
		return nil, err
	}

	return &pb.PingResponse{}, nil
}
