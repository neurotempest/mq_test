package ops

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
)

func StartLoops(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	pingProducerForever(ctx, wg, s)
}

func pingProducerForever(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.After(10*time.Second):
				err := s.GetProducerClient().Ping(ctx, "hello from consumer")
				if err != nil {
					log.Println("ping err:", err.Error())
				}
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()
}
