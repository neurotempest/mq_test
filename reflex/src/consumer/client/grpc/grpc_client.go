package grpc

import (
	"context"
	"flag"

	"github.com/neurotempest/mq_test/reflex/src/consumer"
	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
	"github.com/neurotempest/mq_test/reflex/src/consumer/server"
)

var address = flag.String("consumer_grpc_address", "", "host:port of consumer gRPC service")

var _ consumer.Client = (*client)(nil)

type client struct {
	rpcConn *grpc.ClientConn
	rpcClient server.ProducerClient
}

func New() (*client, error) {
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		if conn.GetState() == connectivity.Ready {
			break
		}
		if !conn.WaitForStateChange(ctx, conn.GetState()) {
			return nil, errors.New("grpc timeout whilst connecting")
		}
	}

	return &client{
		rpcConn: conn,
		rpcClient: userspb.NewUsersClient(conn),
	}, nil
}

func (c *client) Ping(
	ctx contex.Context,
	msg string,
) error {

	_, err := c.rpcClient.Ping(
		ctx,
		&server.PingRequest{
			msg,
		},
	)
	return err
}
