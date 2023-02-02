package localcache

import (
	"errors"
	"time"

	cache "github.com/patrickmn/go-cache"
)

type LocalCache struct {
	data *cache.Cache
}

func NewCache() (lc LocalCache) {
	lc = LocalCache{}
	lc.data = cache.New(5*time.Minute, 10*time.Minute)
	return lc
}

func (mc *LocalCache) Set(id string, item interface{}) error {
	mc.data.Set(id, item, cache.DefaultExpiration)
	return nil
}

func (mc LocalCache) Get(id string) (interface{}, error) {
	item, found := mc.data.Get(id)
	if found {
		return item, nil
	}
	return item, errors.New("ID not found")
}
