// +build !confonly

package ratelimit

//go:generate errorgen

import (
	jrl "github.com/juju/ratelimit"
	"sync"
	"v2ray.com/core/features/ratelimit"
)

type Limiter struct {
	rate     float64
	capacity int64
}

func (l *Limiter) Value() (float64, int64) {
	return l.rate, l.capacity
}

func (l *Limiter) Set(rate float64, capacity int64) {
	l.rate = rate
	l.capacity = capacity
}

func (l *Limiter) Clear() {
	l.rate = 0
	l.capacity = 0
}

type Manager struct {
	access  sync.RWMutex
	buckets map[string]*jrl.Bucket
}

func (*Manager) Type() interface{} {
	return ratelimit.ManagerType()
}

func (m *Manager) RegisterCounter(name string, limiter Limiter) (*jrl.Bucket, error) {
	m.access.Lock()
	defer m.access.Unlock()

	if _, found := m.buckets[name]; found {
		return nil, newError("Bucket ", name, " already registered.")
	}
	newError("create new counter ", name).AtDebug().WriteToLog()
	bucket := jrl.NewBucketWithRate(limiter.rate, limiter.capacity)
	m.buckets[name] = bucket
	return bucket, nil
}

func (m *Manager) GetCounter(name string) *jrl.Bucket {
	m.access.RLock()
	defer m.access.RUnlock()

	if c, found := m.buckets[name]; found {
		return c
	}
	return nil
}

// Start implements common.Runnable.
func (m *Manager) Start() error {
	return nil
}

// Close implement common.Closable.
func (m *Manager) Close() error {
	return nil
}
