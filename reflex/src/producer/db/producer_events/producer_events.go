package producer_events

import (
	"context"
	"database/sql"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"

	"github.com/neurotempest/mq_test/reflex/src/producer"
)

var events = rsql.NewEventsTableInt("producer_events")

func ToStream(dbc *sql.DB) reflex.StreamFunc {
	return events.ToStream(dbc)
}

func FillGaps(dbc *sql.DB) {
	rsql.FillGaps(dbc, events)
}

func Insert(
	ctx context.Context,
	db *sql.DB,
	eventType producer.EventType,
	foreignID int64,
) error {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = events.Insert(ctx, tx, foreignID, eventType)
	if err != nil {
		return err
	}

	return tx.Commit()
}
