package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) SAdd(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SAdd(ctx, key, values...)
	default:
		return redis.Client.SAdd(ctx, key, values...)
	}
}

func (redis RedisClient) SMembers(ctx context.Context, key string) *rds.StringSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SMembers(ctx, key)
	default:
		return redis.Client.SMembers(ctx, key)
	}
}

func (redis RedisClient) SRem(ctx context.Context, key string, members ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SRem(ctx, key, members...)
	default:
		return redis.Client.SRem(ctx, key, members...)
	}
}

func (redis RedisClient) SIsMember(ctx context.Context, key string, members interface{}) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SIsMember(ctx, key, members)
	default:
		return redis.Client.SIsMember(ctx, key, members)
	}
}

func (redis RedisClient) SRandMember(ctx context.Context, key string) *rds.StringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SRandMember(ctx, key)
	default:
		return redis.Client.SRandMember(ctx, key)
	}
}

func (redis RedisClient) SCard(ctx context.Context, key string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.SCard(ctx, key)
	default:
		return redis.Client.SCard(ctx, key)
	}
}
