package lock

import (
	"sync"
	"testing"
)

func BenchmarkLock(b *testing.B) {
	c := Conn{
		data: map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
			"d": "4",
		},
	}

	b.Run("get one", func(b *testing.B) {
		out := make(map[string]string, 4)
		out["a"] = c.GetOne("a")
		out["b"] = c.GetOne("b")
		out["c"] = c.GetOne("c")
		out["d"] = c.GetOne("d")
	})

	b.Run("get batch", func(b *testing.B) {
		c.GetBatch("a", "b", "c", "d")
	})
}

type Conn struct {
	sync.RWMutex
	data map[string]string
}

func (c *Conn) GetOne(key string) string {
	c.RLock()
	defer c.RUnlock()
	v, _ := c.data[key]
	return v
}

func (c *Conn) GetBatch(keys ...string) map[string]string {
	c.RLock()
	defer c.RUnlock()
	out := make(map[string]string, len(keys))
	for _, key := range keys {
		v, _ := c.data[key]
		out[key] = v
	}
	return out
}
