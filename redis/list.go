package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// 队列

func (redis RedisClient) LPop(ctx context.Context, key string) *rds.StringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.LPop(ctx, key)
	default:
		return redis.Client.LPop(ctx, key)
	}
}

func (redis RedisClient) RPop(ctx context.Context, key string) *rds.StringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.RPop(ctx, key)
	default:
		return redis.Client.RPop(ctx, key)
	}
}

func (redis RedisClient) RPush(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.RPush(ctx, key, values)
	default:
		return redis.Client.RPush(ctx, key, values)
	}
}

func (redis RedisClient) LPush(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.LPush(ctx, key, values)
	default:
		return redis.Client.LPush(ctx, key, values)
	}
}

func (redis RedisClient) LTrim(ctx context.Context, key string, start, stop int64) *rds.StatusCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.LTrim(ctx, key, start, stop)
	default:
		return redis.Client.LTrim(ctx, key, start, stop)
	}
}
