package persistence

import (
	"github.com/hashicorp/golang-lru"
	"time"
)

type item struct {
	createdAt time.Time
	value     interface{}
}

type LruCache struct {
	c           *lru.Cache
	maxDuration time.Duration
}

func NewLruCache(maxEntries int, maxDuration time.Duration) (c LruCache, err error) {
	c0, err := lru.New(maxEntries)
	if err != nil {
		return
	}

	c = LruCache{
		c:           c0,
		maxDuration: maxDuration,
	}

	return
}

func (c LruCache) Add(key string, value interface{}) {
	c.c.Add(key, item{createdAt: time.Now(), value: value})
}

func (c LruCache) Get(key string) (interface{}, bool) {
	v, ok := c.c.Get(key)
	if !ok {
		return nil, false
	}

	i, ok := v.(item)
	if !ok {
		return nil, false
	}

	if time.Now().Sub(i.createdAt) > c.maxDuration {
		return nil, false
	}

	return i.value, true
}
