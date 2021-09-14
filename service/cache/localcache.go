package localcache

import (
	"github.com/fanjindong/go-cache"
	"time"
)

var goCache cache.ICache

func init(){
	goCache = cache.NewMemCache(cache.WithShards(8),cache.WithClearInterval(5*time.Minute))
}

func SetCacheE(k string, x interface{}, d time.Duration) {
	goCache.Set(k, x, cache.WithEx(d))
}

func SetCache(k string, x interface{}) {
	goCache.Set(k, x)
}

func GetCache(k string)(interface{},bool){
	return goCache.Get(k)
}

