package service

import (
	"memoryCache/pkg/cache"
	"time"
)

type MemoryCache interface {
	Set(string, int64, time.Duration)
	Get(string) (int64, error)
	Search(int64) ([]string, error)
	Delete(string) error
	Increase(string, int64) error
	Reduce(string, int64) error
}

type Service struct {
	MemoryCache
}

func NewService(cache *cache.Cache) *Service {
	return &Service{
		MemoryCache: NewMemoryCacheService(cache),
	}
}
