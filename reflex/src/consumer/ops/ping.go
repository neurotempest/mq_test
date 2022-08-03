package ops

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/neurotempest/mq_test/reflex/src/consumer/state"
)

func Ping(
	ctx context.Context,
	s state.State,
	msg string,
) error {

	if strings.Contains(msg, "ERROR_PING") {
		return fmt.Errorf("consumer erroring from ping:", msg)
	}

	log.Println("consumer recived ping:", msg)
	return nil
}

