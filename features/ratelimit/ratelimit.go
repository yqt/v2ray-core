package ratelimit

import (
	jrl "github.com/juju/ratelimit"
	"v2ray.com/core/features"
)

//go:generate errorgen

type Limiter interface {
	Value() (float64, int64)
	Set(rate float64, capacity int64)
	Clear()
}

type Manager interface {
	features.Feature

	RegisterBucket(string, Limiter) (*jrl.Bucket, error)
	GetBucket(string) *jrl.Bucket
}

func GetOrRegisterBucket(m Manager, name string, limiter Limiter) (*jrl.Bucket, error) {
	bucket := m.GetBucket(name)
	if bucket != nil {
		return bucket, nil
	}

	return m.RegisterBucket(name, limiter)
}

func ManagerType() interface{} {
	return (*Manager)(nil)
}
