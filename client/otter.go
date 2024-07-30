package client

import (
	"time"

	"github.com/maypok86/otter"
)

type Otter[K comparable, V any] struct {
	client otter.Cache[K, V]
}

func (c *Otter[K, V]) Init(capacity int) {
	client, err := otter.MustBuilder[K, V](capacity).WithTTL(time.Hour).Build()
	if err != nil {
		panic(err)
	}
	c.client = client
}

func (c *Otter[K, V]) Name() string {
	return "otter"
}

func (c *Otter[K, V]) Get(key K) (V, bool) {
	return c.client.Get(key)
}

func (c *Otter[K, V]) Set(key K, value V) {
	c.client.Set(key, value)
}

func (c *Otter[K, V]) Close() {
	c.client.Close()
	c.client = otter.Cache[K, V]{}
}
