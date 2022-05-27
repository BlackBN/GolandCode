package lru

import "container/list"

type Value interface {
	//返回 Value 值的大小
	Len() int
}

type entry struct {
	//保存一个 key , 方便从 cache 里删除
	key   string
	value Value
}

type lruCache struct {
	maxBytes  int64
	userBytes int64
	ll        *list.List
	cache     map[string]*list.Element
	onEvicted func(string, Value)
}

func New(maxBytes int64, onEvicted func(string, Value)) *lruCache {
	return &lruCache{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
	}
}

func (l *lruCache) Get(key string) (value Value, ok bool) {
	if e, ok := l.cache[key]; ok {
		l.ll.MoveToBack(e)
		oldEntry := e.Value.(*entry)
		return oldEntry.value, true
	}
	return
}

func (l *lruCache) Add(key string, val Value) {
	if e, ok := l.cache[key]; ok {
		l.ll.MoveToBack(e)
		oldEntry := e.Value.(*entry)
		l.userBytes -= int64(val.Len() - oldEntry.value.Len())
		oldEntry.value = val
	} else {
		ele := l.ll.PushBack(&entry{key: key, value: val})
		l.cache[key] = ele
		l.userBytes += int64(val.Len() + len(key))

	}
	for l.maxBytes != 0 && l.maxBytes < l.userBytes {
		l.Remove()
	}
}

func (l *lruCache) Remove() {
	if e := l.ll.Front(); e != nil {
		l.ll.Remove(e)
		deleteEntry := e.Value.(*entry)
		delete(l.cache, deleteEntry.key)
		l.userBytes -= int64(deleteEntry.value.Len() + len(deleteEntry.key))
		if l.onEvicted != nil {
			l.onEvicted(deleteEntry.key, deleteEntry.value)
		}
	}
}
