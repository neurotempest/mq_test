package ops

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/neurotempest/mq_test/reflex/src/producer/state"
)

func StartLoops(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	pingConsumerForever(ctx, wg, s)
}

func pingConsumerForever(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.After(10*time.Second):
				err := s.GetConsumerClient().Ping(ctx, "hello from producer")
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
