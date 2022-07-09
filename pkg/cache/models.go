package cache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]ItemCache
}

type ItemCache struct {
	Value      int64
	Created    time.Time
	Expiration int64
}
