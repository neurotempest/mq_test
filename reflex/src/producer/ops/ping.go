package ops

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/neurotempest/mq_test/reflex/src/producer/state"
)

func Ping(
	ctx context.Context,
	s state.State,
	msg string,
) error {

	if strings.Contains(msg, "ERROR_PING") {
		return fmt.Errorf("producer erroring from ping:", msg)
	}

	log.Println("producer recived ping:", msg)
	return nil
}
