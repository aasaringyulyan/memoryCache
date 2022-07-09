package cache

import (
	"errors"
	"time"
)

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]ItemCache)

	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	if cleanupInterval > 0 {
		cache.StartGC()
	}

	return &cache
}

func (c *Cache) Set(key string, value int64, duration time.Duration) {
	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.items[key] = ItemCache{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}

func (c *Cache) Get(key string) (int64, error) {
	item, found := c.items[key]
	if !found {
		return 0, errors.New("key not found")
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return 0, errors.New("cache is outdated")
		}
	}

	return item.Value, nil
}

func (c *Cache) Delete(key string) error {
	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}

func (c *Cache) Search(number int64) ([]string, error) {
	keys := make([]string, 0)

	for index, value := range c.items {
		if value.Value == number {
			if value.Expiration > 0 {
				if time.Now().UnixNano() > value.Expiration {
					break
				}
			}

			keys = append(keys, index)
		}
	}

	if len(keys) == 0 {
		return nil, errors.New("value not found")
	}

	return keys, nil
}

func (c *Cache) Increase(key string, N int64) error {
	item, found := c.items[key]
	if !found {
		return errors.New("key not found")
	} else {

	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return errors.New("cache is outdated")
		}
	}

	item.Value = item.Value + N
	c.items[key] = item

	return nil
}

func (c *Cache) Reduce(key string, N int64) error {
	item, found := c.items[key]
	if !found {
		return errors.New("key not found")
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return errors.New("cache is outdated")
		}
	}

	item.Value = item.Value - N
	c.items[key] = item

	return nil
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {
	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

func (c *Cache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) clearItems(keys []string) {
	c.Lock()
	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
