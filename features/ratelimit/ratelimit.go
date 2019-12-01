package ratelimit

import (
	jrl "github.com/juju/ratelimit"
	"v2ray.com/core/features"
)

//go:generate errorgen

type Limiter interface {
	Value() (float64, int64)
	Set(rate float64, capacity int64)
	GetBucket() *jrl.Bucket
	Clear()
}

type Manager interface {
	features.Feature

	RegisterLimiter(string, float64, int64) (Limiter, error)
	GetLimiter(string) Limiter
}

func GetLimiter(m Manager, name string) Limiter {
	return m.GetLimiter(name)
}

func GetOrRegisterLimiter(m Manager, name string, rate float64, capacity int64) (Limiter, error) {
	limiter := m.GetLimiter(name)
	if limiter != nil {
		return limiter, nil
	}

	return m.RegisterLimiter(name, rate, capacity)
}

func ManagerType() interface{} {
	return (*Manager)(nil)
}

// NoopManager is an implementation of Manager, which doesn't has actual functionalities.
type NoopManager struct{}

// Type implements common.HasType.
func (NoopManager) Type() interface{} {
	return ManagerType()
}

// RegisterCounter implements Manager.
func (NoopManager) RegisterLimiter(string, float64, int64) (Limiter, error) {
	return nil, newError("not implemented")
}

// GetCounter implements Manager.
func (NoopManager) GetLimiter(string) Limiter {
	return nil
}

// Start implements common.Runnable.
func (NoopManager) Start() error { return nil }

// Close implements common.Closable.
func (NoopManager) Close() error { return nil }
