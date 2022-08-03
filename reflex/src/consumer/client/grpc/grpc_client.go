package grpc

import (
	"context"
	"flag"

	"google.golang.org/grpc"

	"github.com/neurotempest/mq_test/reflex/src/consumer"
	"github.com/neurotempest/mq_test/reflex/src/consumer/pb"
)

var address = flag.String("consumer_grpc_address", "", "host:port of consumer gRPC service")

var _ consumer.Client = (*client)(nil)

type client struct {
	rpcConn *grpc.ClientConn
	rpcClient pb.ConsumerClient
}

func New() (*client, error) {
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &client{
		rpcConn: conn,
		rpcClient: pb.NewConsumerClient(conn),
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
