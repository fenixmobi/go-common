package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

func (redis RedisClient) Publish(ctx context.Context, channel string, message interface{}) *rds.IntCmd {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Publish(ctx, channel, message)
	default:
		return redis.Client.Publish(ctx, channel, message)
	}
}

func (redis RedisClient) Subscribe(ctx context.Context, channels ...string) *rds.PubSub {
	switch redis.mode {
	case modeCluster:
		return redis.Cluster.Subscribe(ctx, channels...)
	default:
		return redis.Client.Subscribe(ctx, channels...)
	}
}
