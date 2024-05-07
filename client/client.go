package client

import (
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
)

type CircuitBreakerProxy struct {
	Logger zap.Logger
	Gb     *gobreaker.CircuitBreaker
}
