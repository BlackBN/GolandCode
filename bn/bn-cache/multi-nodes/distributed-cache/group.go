package cache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (gf GetterFunc) Get(key string) ([]byte, error) {
	return gf(key)
}

type Group struct {
	name      string
	getter    Getter
	syncCache *syncCache
	peers     PeerPicker
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, getterFunc GetterFunc, cacheBytes int64) *Group {

	if getterFunc == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getterFunc,
		syncCache: &syncCache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	//如果key为空则查不到
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	// 查询本地缓存，查到就返回，查的是 lruCache
	if bv, ok := g.syncCache.get(key); ok {
		log.Println("[cache] hit !!!")
		return bv, nil
	}
	// 本地查询如果查不到，则查询远程节点的缓存，即实现了分布式缓存的节点
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	//判断是否有远程节点，如果有则需要根据key按照一致性hash算法获取到节点信息，本案例为ip地址
	if g.peers != nil {
		//调用PickPeer方法，获取到真实节点信息
		if peer, ok := g.peers.PickPeer(key); ok {

			if value, err = g.getFromPeer(peer, key); err == nil {
				return value, nil
			}
			log.Println("[GeeCache] Failed to get from peer", err)
		}
	}

	return g.getLocally(key)
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

func (g *Group) getLocally(key string) (ByteView, error) {
	b, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, fmt.Errorf("cannot get value from key %s", key)
	}
	bv := ByteView{b: cloneBytes(b)}
	g.populateCache(key, bv)
	return bv, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.syncCache.add(key, value)
}

//注册
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}
