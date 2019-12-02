// +build !confonly

package ratelimit

//go:generate errorgen

import (
	"context"
	jrl "github.com/juju/ratelimit"
	"sync"
	"v2ray.com/core/features/ratelimit"
)

type Limiter struct {
	rate     float64
	capacity int64
	bucket   *jrl.Bucket
}

func (l *Limiter) Value() (float64, int64) {
	return l.rate, l.capacity
}

func (l *Limiter) Set(rate float64, capacity int64) {
	l.rate = rate
	l.capacity = capacity
	if rate > 0 && capacity > 0 {
		l.bucket = jrl.NewBucketWithRate(rate, capacity)
	}
}

func (l *Limiter) GetBucket() *jrl.Bucket {
	return l.bucket
}

func (l *Limiter) Clear() {
	l.rate = 0
	l.capacity = 0
	l.bucket = nil
}

type limitSetting struct {
	rate     float64
	capacity int64
}

type Manager struct {
	access   sync.RWMutex
	setting  map[string]*limitSetting
	limiters map[string]*Limiter
}

func NewManager(ctx context.Context, config *Config) (*Manager, error) {
	m := &Manager{
		limiters: make(map[string]*Limiter),
		setting:  make(map[string]*limitSetting),
	}

	for _, rule := range config.GetRule() {
		settings := rule.GetSettings()
		upRate := settings.GetUpRate()
		upCapacity := settings.GetUpCapacity()
		if upCapacity == 0 && upRate != 0 {
			upCapacity = upRate * 2
		}
		downRate := settings.GetDownRate()
		downCapacity := settings.GetDownCapacity()
		if downCapacity == 0 && downRate != 0 {
			downCapacity = downRate * 2
		}
		if len(rule.GetUserEmail()) != 0 {
			for _, userEmail := range rule.GetUserEmail() {
				if len(userEmail) == 0 {
					continue
				}
				limiterName := "user:" + userEmail + ":uplink"
				m.addSetting(limiterName, float64(upRate), upCapacity)

				limiterName = "user:" + userEmail + ":downlink"
				m.addSetting(limiterName, float64(downRate), downCapacity)
			}
		}
		if len(rule.GetInboundTag()) != 0 {
			for _, inboundTag := range rule.GetInboundTag() {
				if len(inboundTag) == 0 {
					continue
				}
				if upRate > 0 {
					limiterName := "inboundTag:" + inboundTag + ":uplink"
					m.addSetting(limiterName, float64(upRate), upCapacity)
				}
				if downRate > 0 {
					limiterName := "inboundTag:" + inboundTag + ":downlink"
					m.addSetting(limiterName, float64(downRate), downCapacity)
				}
			}
		}
	}

	return m, nil
}

func (*Manager) Type() interface{} {
	return ratelimit.ManagerType()
}

func (m *Manager) addSetting(name string, rate float64, capacity int64) error {
	if _, found := m.setting[name]; found {
		return newError("setting ", name, " already existed.")
	}
	ls := &limitSetting{rate: rate, capacity: capacity}
	m.setting[name] = ls
	return nil
}

func (m *Manager) RegisterLimiter(name string) (ratelimit.Limiter, error) {
	m.access.Lock()
	defer m.access.Unlock()

	if _, found := m.limiters[name]; found {
		return nil, newError("Limiter ", name, " already registered.")
	}
	ls, found := m.setting[name]
	if !found {
		return nil, newError("Limiter ", name, " has no related setting.")
	}
	newError("create new limiter ", name).AtDebug().WriteToLog()
	limiter := new(Limiter)
	limiter.Set(ls.rate, ls.capacity)
	m.limiters[name] = limiter
	return limiter, nil
}

func (m *Manager) GetLimiter(name string) ratelimit.Limiter {
	m.access.RLock()
	defer m.access.RUnlock()

	if l, found := m.limiters[name]; found {
		return l
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
