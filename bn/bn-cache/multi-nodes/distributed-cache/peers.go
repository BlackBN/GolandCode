package cache

import "GolandCode/bn/bn-cache/multi-nodes/protobuf"

// type Peer interface {
// 	Get(group string, key string) ([]byte, error)
// }

// use protobuf
type Peer interface {
	Get(req *protobuf.CacheRequest, resp *protobuf.CacheResponse) error
}

type PickerPeer interface {
	Pick(key string) (peer Peer, ok bool)
}
