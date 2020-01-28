package states_test

import (
	"context"
	"testing"

	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/process"
	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/states"
	"github.com/stretchr/testify/assert"
)

func TestShowIframe_IsFinal(t *testing.T) {
	s := states.ShowIframe{}
	assert.False(t, s.IsFinal())
}

func TestShowIframe_Name(t *testing.T) {
	s := states.ShowIframe{}
	assert.Equal(t, "ShowIframe", s.Name())
}

func TestShowIframe_Rollback(t *testing.T) {
	s := states.ShowIframe{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestShowIframe_Run(t *testing.T) {
	s := states.ShowIframe{}
	p := &process.Process{}

	s.Run(context.Background(), p, nil)

	assert.Equal(t, states.ValidatePayment{}.Name(), p.Context().CurrentStateName, "Next state should be ValidatePayment.")

	assert.Equal(t, s.Run(context.Background(), &process.Process{}, nil), process.RunResult{})
}
