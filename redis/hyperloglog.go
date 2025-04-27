package redis

import (
	"context"
	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) PFAdd(ctx context.Context, key string, elements ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.PFAdd(ctx, key, elements...)
	default:
		return redis.Client.PFAdd(ctx, key, elements...)
	}
}

func (redis RedisClient) PFCount(ctx context.Context, keys ...string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.PFCount(ctx, keys...)
	default:
		return redis.Client.PFCount(ctx, keys...)
	}
}
