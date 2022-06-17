package cache

import "time"

type record struct {
	key        string
	value      string
	isExpired  bool
	expireTime time.Time
}

func (r *record) expire() (ie bool) {
	t := time.Now()
	r.isExpired = t.After(r.expireTime)
	return
}

type Cache struct {
	Record []record
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	for _, r := range c.Record {
		if r.key == key && !r.expire() {
			return r.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	r := record{key: key, value: value, isExpired: false, expireTime: time.Time{}}
	c.Record = append(c.Record, r)
}

func (c *Cache) Keys() []string {
	rs := []string{}
	for _, r := range c.Record {
		r.expire()
		if !r.expire() {
			rs = append(rs, r.key)
		}
	}
	return rs
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	r := record{key: key, value: value, isExpired: false, expireTime: deadline}
	r.expire()
	c.Record = append(c.Record, r)
}
