package cache

import (
	"../../config"
	"../../connections/redis"
)

var CacheStore *redis.ClientType

func init() {
	CacheStore = &redis.Client
	CacheStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_CACHE_DB)
}

func GetStore() *redis.ClientType {
	CacheStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_CACHE_DB)
	return CacheStore
}