package cache

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var CookieCache = New(60)

type Item struct {
	val       []*http.Cookie
	createdAt time.Time
	life      time.Duration
}

// item 是否过期
func (i *Item) isExpire() bool {
	if i.life == 0 {
		return false
	}
	return time.Now().Sub(i.createdAt) > i.life
}

type Cache struct {
	sync.RWMutex
	items    map[string]*Item
	interval int // 多长时间执行一次GC操作单位秒
}

// 新建一个Cache对象 并开启GC
func New(interval int) *Cache {
	c := &Cache{
		items:    make(map[string]*Item),
		interval: interval,
	}
	go c.GC()
	return c
}

func (c *Cache) Get(key string) []*http.Cookie {
	c.RLock()
	defer c.RUnlock()
	if value, ok := c.items[key]; ok {
		if value.isExpire() {
			return nil
		}
		return value.val
	}
	return nil
}

func (c *Cache) Put(key string, value []*http.Cookie, life time.Duration) error {
	c.Lock()
	defer c.Unlock()
	c.items[key] = &Item{
		val:       value,
		createdAt: time.Now(),
		life:      life * time.Second,
	}
	return nil
}

func (c *Cache) Delete(name string) error {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.items[name]; !ok {
		return errors.New("key not exist")
	}
	delete(c.items, name)
	return nil
}

func (c *Cache) Items() map[string]*Item {
	return c.items
}

// 启动一个每隔 x 秒的一次清除过期的操作
func (c *Cache) GC() {
	for {
		select {
		case <-time.After(time.Duration(c.interval) * time.Second):
			if keys := c.GetExpiredKeys(); len(keys) != 0 {
				for _, key := range keys {
					c.Delete(key)
				}
			}
			fmt.Println("GC")
		}
	}
}

//删除过期的key
func (c *Cache) DeleteExpired(keys []string) {
	c.Lock()
	defer c.Unlock()
	for _, key := range keys {
		delete(c.items, key)
	}
}

// 获取所有的key
func (c *Cache) GetExpiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()
	for key, item := range c.items {
		if item.isExpire() {
			keys = append(keys, key)
		}
	}
	return keys
}
