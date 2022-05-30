package hash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           //可以自定 hash 函数
	replicas int            //为了防止数据倾斜，增加虚拟节点的概念，该值表示一台真实节点对应多少台虚拟节点
	keys     []int          //保存了已存在节点对应的hash值，长度为真实节点个数 * replicas
	hashMap  map[int]string //保存了hash值和真实节点名称的映射，一台真实节点有 replicas 个 hash 值， 长度为真实节点个数 * replicas
}

func NewHash(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	// 默认使用 crc32.ChecksumIEEE
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 参数 keys 为主机节点名称、ip等
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			fmt.Printf("actual key %s hash is %d\n", key, hash)
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
	//fmt.Printf("keys : %v\n, hashMap %v\n", m.keys, m.hashMap)
}

// 参数key为需要查找的缓存的key，与Add方法的key不一样概念
func (m *Map) Get(key string) string {
	if m.keys == nil || len(m.keys) == 0 {
		return ""
	}
	// 计算缓存key的hash值
	hash := int(m.hash([]byte(key)))
	// 返回 m.keys 大于等于 hash 值的数组下标
	// idx 为 0~len(m.keys), 当查找不到 m.keys[i] >= hash 的时候返回数组的长度，而不是 -1
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// 正常如果能找到 m.keys[i] >= hash 的话 idx % len(m.keys) == idx
	// 如果不能找到，则 idx == 0  这样刚好形成一个环
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
