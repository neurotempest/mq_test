package ops

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"

	db_cursors "github.com/neurotempest/mq_test/reflex/src/consumer/db/cursors"
	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
)

func StartLoops(
	ctx context.Context,
	wg *sync.WaitGroup,
	s state.State,
) {

	pingProducerForever(ctx, wg, s)
	consumeProducerEventsForever(ctx, s)
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

func consumeProducerEventsForever(
	ctx context.Context,
	s state.State,
) {

	go func() {

		// Wait afew seconds before starting the reflex consumer;
		// otherwise it tries to grab events immediately and fails
		// (if the `producer` service is stood up at the same time)
		// triggering the 1 minute error backoff, meaning we don't
		// start consuming events for 1 minute.
		//select {
		//	case <-time.After(5*time.Second):
		//}

		log.Println("starting consumer...")

		spec := reflex.NewSpec(
			s.GetProducerClient().StreamProducerEvents,
			db_cursors.ToStore(s.GetDb()),
			reflex.NewConsumer(
				"producer_event_consumer",
				consumeProducerEvent,
			),
		)

		rpatterns.RunForever(context.Background, spec)
	}()
}

func consumeProducerEvent(
	ctx context.Context,
	f fate.Fate,
	e *reflex.Event,
) error {

	log.Println("consumed producer event: ID", e.ID, "type", e.Type, "foreign ID", e.ForeignID)
	return f.Tempt()
}

