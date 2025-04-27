package redis

import (
	"context"
	"time"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *rds.Cmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Eval(ctx, script, keys, args)
	default:
		return redis.Client.Eval(ctx, script, keys, args)
	}
}

func (redis RedisClient) Del(ctx context.Context, keys ...string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Del(ctx, keys...)
	default:
		return redis.Client.Del(ctx, keys...)
	}
}

func (redis RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Expire(ctx, key, expiration)
	default:
		return redis.Client.Expire(ctx, key, expiration)
	}
}

func (redis RedisClient) Exists(ctx context.Context, keys ...string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Exists(ctx, keys...)
	default:
		return redis.Client.Exists(ctx, keys...)
	}
}

func (redis RedisClient) TTL(ctx context.Context, keys string) *rds.DurationCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.TTL(ctx, keys)
	default:
		return redis.Client.TTL(ctx, keys)
	}
}
