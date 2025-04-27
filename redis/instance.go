package redis

import (
	"context"
	"time"

	rds "github.com/redis/go-redis/v9"
)

// redis三种部署方式
// 1、单体
// 2、集群

const (
	_ = iota
	modeClient
	modeCluster
)

const Nil = rds.Nil

type Options struct {
	Host         []string
	DB           int
	Password     string
	Mode         int `desc:"1单机,2集群"`
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type RedisClient struct {
	Client  *rds.Client
	Cluster *rds.ClusterClient
	mode    int
}

func InitRedis(opt Options) *RedisClient {
	redis := &RedisClient{}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	redis.mode = opt.Mode
	switch opt.Mode {
	case modeCluster:
		redis.Cluster = rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:        opt.Host,
			Password:     opt.Password,
			WriteTimeout: opt.WriteTimeout,
			ReadTimeout:  opt.ReadTimeout,
		})

		// 检测
		if _, err := redis.Cluster.Ping(ctx).Result(); err != nil {
			panic(err)
		}

	default:
		// 默认单机
		redis.Client = rds.NewClient(&rds.Options{
			Addr:         opt.Host[0],
			Password:     opt.Password,
			DB:           opt.DB,
			WriteTimeout: opt.WriteTimeout,
			ReadTimeout:  opt.ReadTimeout,
		})

		// 检测
		if _, err := redis.Client.Ping(ctx).Result(); err != nil {
			panic(err)
		}
	}

	return redis
}
