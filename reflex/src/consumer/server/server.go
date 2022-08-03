package server

import (
	"context"

	"github.com/neurotempest/mq_test/reflex/src/consumer/ops"
	"github.com/neurotempest/mq_test/reflex/src/consumer/pb"
	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
)

var _ pb.ConsumerServer = (*server)(nil)

type server struct {
	pb.UnimplementedConsumerServer
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
