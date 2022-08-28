package ops

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/neurotempest/mq_test/reflex/src/producer"
	db_producer_events "github.com/neurotempest/mq_test/reflex/src/producer/db/producer_events"
	"github.com/neurotempest/mq_test/reflex/src/producer/state"
)

func StartLoops(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	pingConsumerForever(ctx, wg, s)

	createProducerEventsForever(ctx, wg, s)
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

func createProducerEventsForever(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.After(10*time.Second):
				err := createRandomProducerEvent(ctx, s)
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

func createRandomProducerEvent(
	ctx context.Context,
	s state.State,
) error {

	log.Println("creating producer event")

	return db_producer_events.Insert(
		ctx,
		s.GetDb(),
		// TODO make random
		producer.EventTypeOne,
		1234,
	)
}
