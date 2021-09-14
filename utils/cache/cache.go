package cache

import (
	"github.com/muesli/cache2go"
)
var cache *cache2go.CacheTable

func getCache() *cache2go.CacheTable{
	cache:=cache2go.Cache("go-blog")
	return cache
}
