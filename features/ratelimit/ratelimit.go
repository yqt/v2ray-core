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

	RegisterLimiter(string) (Limiter, error)
	GetLimiter(string) Limiter
}

func RequireLimiter(m Manager, userEmail string, inboundTag string, isUp bool) (Limiter, error) {
	userLimiterName := "user:" + userEmail
	userWildLimiterName := "user:*"
	inboundTagLimiterName := "inboundTag:" + inboundTag
	var limiterNameTail string
	if isUp {
		limiterNameTail = ":uplink"
	} else {
		limiterNameTail = ":downlink"
	}
	userLimiterName += limiterNameTail
	userWildLimiterName += limiterNameTail
	inboundTagLimiterName += limiterNameTail
	if l, err := getOrRegisterLimiter(m, userLimiterName); l != nil {
		return l, err
	}
	if l, err := getOrRegisterLimiter(m, userWildLimiterName); l != nil {
		return l, err
	}
	if l, err := getOrRegisterLimiter(m, inboundTagLimiterName); l != nil {
		return l, err
	}
	return nil, newError("no related limiter found.")
}

func getOrRegisterLimiter(m Manager, name string) (Limiter, error) {
	limiter := m.GetLimiter(name)
	if limiter != nil {
		return limiter, nil
	}

	return m.RegisterLimiter(name)
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
