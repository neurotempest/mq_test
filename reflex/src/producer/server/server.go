package server

import (
	"context"

	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"

	db_producer_events "github.com/neurotempest/mq_test/reflex/src/producer/db/producer_events"
	"github.com/neurotempest/mq_test/reflex/src/producer/ops"
	"github.com/neurotempest/mq_test/reflex/src/producer/pb"
	"github.com/neurotempest/mq_test/reflex/src/producer/state"
)

var _ pb.ProducerServer = (*server)(nil)

type server struct {
	pb.UnimplementedProducerServer
	reflexServer *reflex.Server
	state state.State
}

func New(st state.State) server {

	return server{
		reflexServer: reflex.NewServer(),
		state: st,
	}
}

func (s *server) Stop() {

	s.reflexServer.Stop()
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

func (s *server) StreamProducerEvents(
	req *reflexpb.StreamRequest,
	ss pb.Producer_StreamProducerEventsServer,
) error {

	return s.reflexServer.Stream(
		db_producer_events.ToStream(s.state.GetDb()),
		req,
		ss,
	)
}
