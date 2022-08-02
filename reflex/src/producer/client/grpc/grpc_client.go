package grpc

import (
	"context"
	"errors"
	"flag"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

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
