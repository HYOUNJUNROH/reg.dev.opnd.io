package kvstore

import (
	"context"
	"fmt"
	"os"
	"time"

	"git.dev.opnd.io/gc/backend-admin/pkg/config"
	"git.dev.opnd.io/gc/backend-admin/pkg/logger"
	cache "github.com/SporkHubr/echo-http-cache"
	rediscache "github.com/SporkHubr/echo-http-cache/adapter/redis"
	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
)

var (
	CacheClient *cache.Client
	Store       *redisstore.RedisStore
	RedisClient redis.UniversalClient
)

func Init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			os.Exit(2)
		}
	}()
	var err error
	logger.Logger.Infof("kvstore url : %v", fmt.Sprintf("%v:%v", config.Config.Cache.Host, config.Config.Cache.Port))

	ringOpt := &rediscache.RingOptions{
		Addrs: map[string]string{
			"server1": fmt.Sprintf("%v:%v", config.Config.Cache.Host, config.Config.Cache.Port),
		},
		Password: config.Config.Cache.Password,
	}
	CacheClient, err = cache.NewClient(
		cache.ClientWithAdapter(rediscache.NewAdapter(ringOpt)),
		cache.ClientWithTTL(10*time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	RedisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		MasterName:       config.Config.Cache.MasterName,
		Addrs:            []string{fmt.Sprintf("%v:%v", config.Config.Cache.Host, config.Config.Cache.Port)},
		Password:         config.Config.Cache.Password,
		SentinelPassword: config.Config.Cache.Password,
	})

	Store, err = redisstore.NewRedisStore(context.Background(), RedisClient)
	if err != nil {
		logger.Logger.Info("failed to create redis store by sentinel: ", err)
		RedisClient = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    []string{fmt.Sprintf("%v:%v", config.Config.Cache.Host, config.Config.Cache.Port)},
			Password: config.Config.Cache.Password,
		})

		Store, err = redisstore.NewRedisStore(context.Background(), RedisClient)
		if err != nil {
			logger.Logger.Fatal("failed to create redis store: ", err)
		}
	}
}
