package state

//go:generate stategen --inputFile=state.go --inputStruct=stateImpl --outputInterface=State --outputFile=state_gen.go

import (
	"log"

	"github.com/neurotempest/mq_test/reflex/src/producer"
	producer_grpc "github.com/neurotempest/mq_test/reflex/src/producer/client/grpc"
)

type stateImpl struct {
	producerClient producer.Client
}

func New() *stateImpl {

	producerClient, err := producer_grpc.New()
	if err != nil {
		log.Fatal("failed to create producer grpc client:", err.Error())
	}

	return &stateImpl{
		producerClient: producerClient,
	}
}

