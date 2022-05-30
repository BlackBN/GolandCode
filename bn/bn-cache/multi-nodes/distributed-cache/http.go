package cache

import (
	hash "GolandCode/bn/bn-cache/consistent-hash"
	"GolandCode/bn/bn-cache/multi-nodes/protobuf"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
)

const (
	defaultBasePath = "/_cache"
	defaultReplicas = 50
)

type httpPeer struct {
	baseURL string
}

// func (h *httpPeer) Get(group string, key string) ([]byte, error) {
func (h *httpPeer) Get(req *protobuf.CacheRequest, resp *protobuf.CacheResponse) (err error) {
	// u := fmt.Sprintf("%v/%v/%v", h.baseURL, url.QueryEscape(group), url.QueryEscape(key))
	u := fmt.Sprintf("%v/%v/%v", h.baseURL, url.QueryEscape(req.Group), url.QueryEscape(req.Key))
	fmt.Printf("peer url is : %s\n", u)
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}
	// grpc 反序列化
	if err = proto.Unmarshal(bytes, resp); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}
	return nil
}

var _ Peer = (*httpPeer)(nil)

type httpPool struct {
	self      string
	basePath  string
	mu        sync.Mutex
	hashMap   *hash.Map
	httpPeers map[string]*httpPeer
}

func NewHttpPool(self string) *httpPool {
	return &httpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}
func (h *httpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !strings.HasPrefix(path, h.basePath) {
		panic("HTTPPool serving unexcepted path :" + path)
	}
	h.Log("%s %s", r.Method, path)
	parts := strings.SplitN(path[len(h.basePath)+1:], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	fmt.Printf("groupname : %s , key : %s\n", groupName, key)
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusBadRequest)
		return
	}
	bv, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//use grpc
	body, err := proto.Marshal(&protobuf.CacheResponse{Value: bv.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

func (h *httpPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", h.self, fmt.Sprintf(format, v...))
}

//
func (h *httpPool) Set(addrs ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.hashMap = hash.NewHash(defaultReplicas, nil)
	h.hashMap.Add(addrs...)
	h.httpPeers = make(map[string]*httpPeer, len(addrs))
	for _, addr := range addrs {
		h.httpPeers[addr] = &httpPeer{baseURL: addr + h.basePath}
	}
}

// PickPeer picks a peer according to key
func (h *httpPool) Pick(key string) (Peer, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if peer := h.hashMap.Get(key); peer != "" && peer != h.self {
		h.Log("Pick peer %s", peer)
		return h.httpPeers[peer], true
	}
	return nil, false
}

var _ PickerPeer = (*httpPool)(nil)
