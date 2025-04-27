package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) *rds.FloatCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZIncrBy(ctx, key, increment, member)
	default:
		return redis.Client.ZIncrBy(ctx, key, increment, member)
	}
}

func (redis RedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) *rds.StringSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRevRange(ctx, key, start, stop)
	default:
		return redis.Client.ZRevRange(ctx, key, start, stop)
	}
}

func (redis RedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRem(ctx, key, members...)
	default:
		return redis.Client.ZRem(ctx, key, members...)
	}
}

func (redis RedisClient) ZCard(ctx context.Context, key string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZCard(ctx, key)
	default:
		return redis.Client.ZCard(ctx, key)
	}
}

func (redis RedisClient) ZScore(ctx context.Context, key, member string) *rds.FloatCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZScore(ctx, key, member)
	default:
		return redis.Client.ZScore(ctx, key, member)
	}
}

func (redis RedisClient) ZRank(ctx context.Context, key, member string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRank(ctx, key, member)
	default:
		return redis.Client.ZRank(ctx, key, member)
	}
}

func (redis RedisClient) ZRevRank(ctx context.Context, key, member string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRevRank(ctx, key, member)
	default:
		return redis.Client.ZRevRank(ctx, key, member)
	}
}

func (redis RedisClient) ZRange(ctx context.Context, key string, start, stop int64) *rds.StringSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRange(ctx, key, start, stop)
	default:
		return redis.Client.ZRange(ctx, key, start, stop)
	}
}

func (redis RedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *rds.ZSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRangeWithScores(ctx, key, start, stop)
	default:
		return redis.Client.ZRangeWithScores(ctx, key, start, stop)
	}
}

func (redis RedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *rds.ZSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRevRangeWithScores(ctx, key, start, stop)
	default:
		return redis.Client.ZRevRangeWithScores(ctx, key, start, stop)
	}
}

func (redis RedisClient) ZRangeByScore(ctx context.Context, key string, opt *rds.ZRangeBy) *rds.StringSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRangeByScore(ctx, key, opt)
	default:
		return redis.Client.ZRangeByScore(ctx, key, opt)
	}
}

func (redis RedisClient) ZAdd(ctx context.Context, key string, members ...rds.Z) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZAdd(ctx, key, members...)
	default:
		return redis.Client.ZAdd(ctx, key, members...)
	}
}

// ZRevRangeByScore
func (redis RedisClient) ZRevRangeByScore(ctx context.Context, key string, opt *rds.ZRangeBy) *rds.StringSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZRevRangeByScore(ctx, key, opt)
	default:
		return redis.Client.ZRevRangeByScore(ctx, key, opt)
	}
}

func (redis RedisClient) ZCount(ctx context.Context, key, min, max string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.ZCount(ctx, key, min, max)
	default:
		return redis.Client.ZCount(ctx, key, min, max)
	}
}
