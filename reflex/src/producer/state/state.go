package state

//go:generate stategen --inputFile=state.go --inputStruct=stateImpl --outputInterface=State --outputFile=state_gen.go

import (
	"log"

	"github.com/neurotempest/mq_test/reflex/src/consumer"
	consumer_grpc "github.com/neurotempest/mq_test/reflex/src/consumer/client/grpc"
)

type stateImpl struct {
	consumerClient consumer.Client
}

func New() *stateImpl {

	consumerClient, err := consumer_grpc.New()
	if err != nil {
		log.Fatal("failed to create consumer grpc client:", err.Error())
	}

	return &stateImpl{
		consumerClient: consumerClient,
	}
}

