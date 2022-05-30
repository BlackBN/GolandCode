package cache

import (
	"GolandCode/bn/bn-cache/multi-nodes/distributed-cache/singleflight"
	"GolandCode/bn/bn-cache/multi-nodes/protobuf"
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
	name       string
	getter     Getter
	syncCache  *syncCache
	pickerPeer PickerPeer
	loader     *singleflight.Group
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
		loader:    &singleflight.Group{},
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
		log.Println("[Cache] hit !!!")
		return bv, nil
	}
	// 本地查询如果查不到，则查询远程节点的缓存，即实现了分布式缓存的节点
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	//判断是否有远程节点，如果有则需要根据key按照一致性hash算法获取到节点信息，本案例为ip地址
	val, err := g.loader.Do(key, func() (interface{}, error) {
		if g.pickerPeer != nil {
			//调用PickPeer方法，获取到真实节点信息
			if peer, ok := g.pickerPeer.Pick(key); ok {
				fmt.Printf("get key %s value from peer %s\n", key, peer)
				if value, err = g.getFromPeer(peer, key); err == nil {
					return value, nil
				}
				fmt.Printf("[GeeCache] Failed to get from peer %v\n", err)
			}
		}
		return g.getLocally(key)
	})
	if err == nil {
		return val.(ByteView), err
	}
	return
}

func (g *Group) getFromPeer(peer Peer, key string) (ByteView, error) {
	resp := &protobuf.CacheResponse{}
	err := peer.Get(&protobuf.CacheRequest{Group: g.name, Key: key}, resp)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: resp.Value}, nil
}

func (g *Group) getLocally(key string) (ByteView, error) {
	fmt.Printf("get %s value from db\n", key)
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
func (g *Group) RegisterPeers(pickerPeer PickerPeer) {
	if g.pickerPeer != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.pickerPeer = pickerPeer
}
