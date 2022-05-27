package cache

import (
	"GolandCode/bn/bn-cache/sync-cache/lru"
	"sync"
)

type syncCache struct {
	mu         sync.Mutex
	lruCache   *lru.LruCache
	cacheBytes int64
}

func (sc *syncCache) add(key string, value ByteView) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	if sc.lruCache == nil {
		sc.lruCache = lru.New(sc.cacheBytes, nil)
	}
	sc.lruCache.Add(key, value)
}

func (sc *syncCache) get(key string) (bv ByteView, ok bool) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	if sc.lruCache == nil {
		return
	}
	if val, ok := sc.lruCache.Get(key); ok {
		return val.(ByteView), true
	}
	return
}
