// Code generated by "cmap-gen -package maps -type User"; DO NOT EDIT.

package maps

import "hash/fnv"
import "sync"

// UserMap is sharded concurrent map which key type is string and
// value type is User
type UserMap struct {
	shards []*UserShard
	n      uint32
}

// NewUserMap creates new UserMap with specified shards count
func NewUserMap(nShards int) *UserMap {
	shards := make([]*UserShard, nShards)
	for i := 0; i < nShards; i++ {
		shards[i] = NewUserShard()
	}
	return &UserMap{shards: shards, n: uint32(nShards)}
}

func (c UserMap) hash(s string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	return h.Sum32() % c.n
}

// Get returns the value stored by specified key
func (c UserMap) Get(key string) User {
	return c.shards[c.hash(key)].Get(key)
}

// Set stores the specified value under the specified key
func (c UserMap) Set(key string, value User) {
	c.shards[c.hash(key)].Set(key, value)
}

// UserShard is concurrent map which key type is string and
// value type is User
type UserShard struct {
	mu   sync.RWMutex
	data map[string]User
}

// NewUserShard creates new UserShard
func NewUserShard() *UserShard {
	return &UserShard{
		data: make(map[string]User),
	}
}

// Get returns the value stored by specified key
func (c *UserShard) Get(key string) User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// Set stores the specified value under the specified key
func (c *UserShard) Set(key string, value User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}