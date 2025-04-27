package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) BFAdd(ctx context.Context, key string, element any) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.BFAdd(ctx, key, element)
	default:
		return redis.Client.BFAdd(ctx, key, element)
	}
}

func (redis RedisClient) BFExists(ctx context.Context, key string, element any) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.BFExists(ctx, key, element)
	default:
		return redis.Client.BFExists(ctx, key, element)
	}
}

func (redis RedisClient) BFReserveWithArgs(ctx context.Context, key string, options *rds.BFReserveOptions) *rds.StatusCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.BFReserveWithArgs(ctx, key, options)
	default:
		return redis.Client.BFReserveWithArgs(ctx, key, options)
	}
}

func (redis RedisClient) BFCard(ctx context.Context, key string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.BFCard(ctx, key)
	default:
		return redis.Client.BFCard(ctx, key)
	}
}
