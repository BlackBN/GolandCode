package cache

import (
	hash "GolandCode/bn/bn-cache/consistent-hash"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_cache"
	defaultReplicas = 50
)

type httpGetter struct {
	baseURL string
}

func (h *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf("%v%v/%v", h.baseURL, url.QueryEscape(group), url.QueryEscape(key))
	fmt.Printf("httpGetter url is : %s\n", u)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}
	return bytes, nil
}

var _ PeerGetter = (*httpGetter)(nil)

type httpPool struct {
	self        string
	basePath    string
	mu          sync.Mutex
	hashMap     *hash.Map
	httpGetters map[string]*httpGetter
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
	fmt.Printf("groupname : %s , key : %s", groupName, key)
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusBadRequest)
		return
	}
	bv, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(bv.ByteSlice())
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
	h.httpGetters = make(map[string]*httpGetter, len(addrs))
	for _, addr := range addrs {
		h.httpGetters[addr] = &httpGetter{baseURL: addr + h.basePath}
	}
}

// PickPeer picks a peer according to key
func (h *httpPool) PickPeer(key string) (PeerGetter, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if peer := h.hashMap.Get(key); peer != "" && peer != h.self {
		h.Log("Pick peer %s", peer)
		return h.httpGetters[peer], true
	}
	return nil, false
}

var _ PeerPicker = (*httpPool)(nil)
