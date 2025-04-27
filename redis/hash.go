package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) HSet(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HSet(ctx, key, values...)
	default:
		return redis.Client.HSet(ctx, key, values...)
	}
}

func (redis RedisClient) HSetNX(ctx context.Context, key string, field string, values interface{}) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HSetNX(ctx, key, field, values)
	default:
		return redis.Client.HSetNX(ctx, key, field, values)
	}
}

func (redis RedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HMSet(ctx, key, values...)
	default:
		return redis.Client.HMSet(ctx, key, values...)
	}
}

func (redis RedisClient) HExists(ctx context.Context, key, field string) *rds.BoolCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HExists(ctx, key, field)
	default:
		return redis.Client.HExists(ctx, key, field)
	}
}

func (redis RedisClient) HGetAll(ctx context.Context, key string) *rds.MapStringStringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HGetAll(ctx, key)
	default:
		return redis.Client.HGetAll(ctx, key)
	}
}

func (redis RedisClient) HGet(ctx context.Context, key, field string) *rds.StringCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HGet(ctx, key, field)
	default:
		return redis.Client.HGet(ctx, key, field)
	}
}

func (redis RedisClient) HMGet(ctx context.Context, key string, fields ...string) *rds.SliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HMGet(ctx, key, fields...)
	default:
		return redis.Client.HMGet(ctx, key, fields...)
	}
}

func (redis RedisClient) HDel(ctx context.Context, key string, fields ...string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HDel(ctx, key, fields...)
	default:
		return redis.Client.HDel(ctx, key, fields...)
	}
}

func (redis RedisClient) HLen(ctx context.Context, key string) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HLen(ctx, key)
	default:
		return redis.Client.HLen(ctx, key)
	}
}

func (redis RedisClient) HIncrBy(ctx context.Context, key string, fields string, value int64) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HIncrBy(ctx, key, fields, value)
	default:
		return redis.Client.HIncrBy(ctx, key, fields, value)
	}
}

func (redis RedisClient) HIncrByFloat(ctx context.Context, key string, fields string, value float64) *rds.FloatCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HIncrByFloat(ctx, key, fields, value)
	default:
		return redis.Client.HIncrByFloat(ctx, key, fields, value)
	}
}

func (redis RedisClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *rds.ScanCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HScan(ctx, key, cursor, match, count)
	default:
		return redis.Client.HScan(ctx, key, cursor, match, count)
	}
}

func (redis RedisClient) HKeys(ctx context.Context, key string) *rds.StringSliceCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.HKeys(ctx, key)
	default:
		return redis.Client.HKeys(ctx, key)
	}
}
