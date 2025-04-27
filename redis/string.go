package redis

import (
	"context"
	"time"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *rds.StatusCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Set(ctx, key, value, expiration)
	default:
		return redis.Client.Set(ctx, key, value, expiration)
	}
}

func (redis RedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SetNX(ctx, key, value, expiration)
	default:
		return redis.Client.SetNX(ctx, key, value, expiration)
	}
}

func (redis RedisClient) Get(ctx context.Context, key string) *rds.StringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Get(ctx, key)
	default:
		return redis.Client.Get(ctx, key)
	}
}

func (redis RedisClient) MGet(ctx context.Context, keys ...string) *rds.SliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.MGet(ctx, keys...)
	default:
		return redis.Client.MGet(ctx, keys...)
	}
}

func (redis RedisClient) Incr(ctx context.Context, key string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Incr(ctx, key)
	default:
		return redis.Client.Incr(ctx, key)
	}
}

func (redis RedisClient) IncrBy(ctx context.Context, key string, value int64) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.IncrBy(ctx, key, value)
	default:
		return redis.Client.IncrBy(ctx, key, value)
	}
}

func (redis RedisClient) IncrByFloat(ctx context.Context, key string, value float64) *rds.FloatCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.IncrByFloat(ctx, key, value)
	default:
		return redis.Client.IncrByFloat(ctx, key, value)
	}
}

func (redis RedisClient) GetSet(ctx context.Context, key string, value any) *rds.StringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.GetSet(ctx, key, value)
	default:
		return redis.Client.GetSet(ctx, key, value)
	}
}

func (redis RedisClient) Rename(ctx context.Context, key, newKey string) *rds.StatusCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Rename(ctx, key, newKey)
	default:
		return redis.Client.Rename(ctx, key, newKey)
	}
}
