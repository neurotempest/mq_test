package grpc

import (
	"context"
	"flag"

	"google.golang.org/grpc"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"

	"github.com/neurotempest/mq_test/reflex/src/producer"
	"github.com/neurotempest/mq_test/reflex/src/producer/pb"
)

var address = flag.String("producer_grpc_address", ":1234", "host:port of producer gRPC service")

var _ producer.Client = (*client)(nil)

type client struct {
	rpcConn *grpc.ClientConn
	rpcClient pb.ProducerClient
}

func New() (*client, error) {
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &client{
		rpcConn: conn,
		rpcClient: pb.NewProducerClient(conn),
	}, nil
}

func (c *client) Ping(
	ctx context.Context,
	msg string,
) error {

	_, err := c.rpcClient.Ping(
		ctx,
		&pb.PingRequest{
			Msg: msg,
		},
	)
	return err
}

func (c *client) StreamProducerEvents(
	ctx context.Context,
	after string,
	opts ...reflex.StreamOption,
) (reflex.StreamClient, error) {

	streamFn := reflex.WrapStreamPB(
		func(
			ctx context.Context,
			req *reflexpb.StreamRequest,
		) (reflex.StreamClientPB, error) {
			return c.rpcClient.StreamProducerEvents(ctx, req)
		},
	)

	return streamFn(ctx, after, opts...)
}
