package producer

import (
	"context"

	"github.com/luno/reflex"
)

type Client interface {
	Ping(ctx context.Context, msg string) error

	StreamProducerEvents(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error)
}
