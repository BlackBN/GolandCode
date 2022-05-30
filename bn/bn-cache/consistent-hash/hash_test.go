package hash

import (
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	hash := NewHash(3, func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	})
	hash.Add("2", "4", "6")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for key, value := range testCases {
		if hash.Get(key) != value {
			t.Errorf("Asking for %s, should have yielded %s", key, value)
		}
	}

	hash.Add("8")
	testCases["27"] = "8"
	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

}
