package cache

import (
	"time"
)

type record struct {
	value      string
	isExpired  bool
	expireTime time.Time
}

func (r *record) expire() bool {
	if r.expireTime.IsZero() {
		return false
	}
	t := time.Now()
	r.isExpired = t.After(r.expireTime)
	return r.isExpired
}

type Cache struct {
	Record map[string]record
}

func NewCache() Cache {
	rs := make(map[string]record)
	return Cache{rs}
}

func (c *Cache) Get(key string) (string, bool) {
	for k, r := range c.Record {
		if k == key && !r.expire() {
			return r.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.Record[key] = record{value: value, isExpired: false, expireTime: time.Time{}}
}

func (c *Cache) Keys() []string {
	rs := []string{}
	for k, r := range c.Record {
		r.expire()
		if !r.expire() {
			rs = append(rs, k)
		}
	}
	return rs
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	r := record{value: value, isExpired: false, expireTime: deadline}
	r.expire()
	c.Record[key] = r
}
