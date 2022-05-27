package cache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_cache"

type httpPool struct {
	self     string
	basePath string
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
	bv, err := group.get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(bv.ByteSlice())
}

func (h *httpPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", h.self, fmt.Sprintf(format, v...))
}
