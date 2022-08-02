package state

// Generated by stategen. DO NOT EDIT.

import(
	"github.com/neurotempest/mq_test/reflex/src/producer"
	"testing"
)

type State interface {
	GetProducerClient() producer.Client
}

func (s *stateImpl) GetProducerClient() producer.Client {

	return s.producerClient
}

type stateOption func(*stateImpl)

func NewStateForTesting(
	_ testing.TB,
	opts ...stateOption,
) State {

	var s stateImpl
	for _, opt := range opts {
		opt(&s)
	}
	return &s
}

func WithProducerClient(producerClient producer.Client) stateOption {

	return func(s *stateImpl) {
		s.producerClient = producerClient
	}
}
