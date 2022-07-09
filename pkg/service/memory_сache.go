package service

import (
	"memoryCache/pkg/cache"
	"time"
)

type MemoryCacheService struct {
	cache *cache.Cache
}

func NewMemoryCacheService(cache *cache.Cache) *MemoryCacheService {
	return &MemoryCacheService{
		cache: cache,
	}
}

func (s *MemoryCacheService) Set(key string, value int64, duration time.Duration) {
	s.cache.Lock()
	defer s.cache.Unlock()

	s.cache.Set(key, value, duration)

}

func (s *MemoryCacheService) Get(key string) (int64, error) {
	s.cache.RLock()
	defer s.cache.RUnlock()

	value, err := s.cache.Get(key)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (s *MemoryCacheService) Delete(key string) error {
	s.cache.Lock()
	defer s.cache.Unlock()

	err := s.cache.Delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (s *MemoryCacheService) Search(number int64) ([]string, error) {
	s.cache.RLock()
	defer s.cache.RUnlock()

	keys, err := s.cache.Search(number)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (s *MemoryCacheService) Increase(key string, n int64) error {
	s.cache.Lock()
	defer s.cache.Unlock()

	err := s.cache.Increase(key, n)
	if err != nil {
		return err
	}

	return nil
}

func (s *MemoryCacheService) Reduce(key string, n int64) error {
	s.cache.Lock()
	defer s.cache.Unlock()

	err := s.cache.Reduce(key, n)
	if err != nil {
		return err
	}

	return nil
}
