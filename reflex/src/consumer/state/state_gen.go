package state

// Generated by stategen. DO NOT EDIT.

import(
	"database/sql"
	"github.com/neurotempest/mq_test/reflex/src/producer"
	"testing"
)

type State interface {
	GetDb() *sql.DB
	GetProducerClient() producer.Client
}

func (s *stateImpl) GetDb() *sql.DB {

	return s.db
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

func WithDb(db *sql.DB) stateOption {

	return func(s *stateImpl) {
		s.db = db
	}
}

func WithProducerClient(producerClient producer.Client) stateOption {

	return func(s *stateImpl) {
		s.producerClient = producerClient
	}
}
