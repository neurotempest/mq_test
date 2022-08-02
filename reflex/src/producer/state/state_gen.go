package state

// Generated by stategen. DO NOT EDIT.

import(
	"testing"
)

type State interface {
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