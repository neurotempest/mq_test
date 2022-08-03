package consumer

import (
	"context"
)

type Client interface {
	Ping(ctx context.Context, msg string) error
}
