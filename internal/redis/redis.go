package redis

import (
	"fmt"
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/logger"
	"time"

	"github.com/go-redis/redis"
)

func init() {
	if !config.GetConfig().REDIS.Use {
		panic("Redis not open, please check config")
	}
	GetRedis()
}

type MainRedis struct {
	pClient *redis.Client
}

var mainRedis *MainRedis

func initFromConfig() {
	conf := &config.GetConfig().REDIS.Config
	var client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.IP, conf.PORT),
		Password: conf.PASSWORD,
		DB:       conf.DB,
	})
	mainRedis = &MainRedis{
		pClient: client,
	}
	_, err := client.Ping().Result()
	if err != nil {
		logger.Error.Fatalln(err)
		return
	}
	logger.Info.Printf("redis init finish")
}

func GetRedis() *MainRedis {
	if mainRedis == nil {
		initFromConfig()
	}
	return mainRedis
}

// Get emapty result throw error
func (r *MainRedis) Get(key string) (string, error) {
	if r.pClient == nil {
		logger.Error.Fatalln("pClient cannot be nil")
	}
	return r.pClient.Get(key).Result()
}

func (r *MainRedis) GetInt(key string) (int, error) {
	if r.pClient == nil {
		logger.Error.Fatalln("pClient cannot be nil")
	}
	logger.Info.Printf("get val of key:%s", key)
	return r.pClient.Get(key).Int()
}

func (r *MainRedis) GetIntOrDefault(key string, def int) (int, error) {
	val, err := r.GetInt(key)
	if err == nil {
		return val, err
	}
	if err == redis.Nil {
		return def, nil //已处理redis.Nil异常
	}
	return def, err
}

func (r *MainRedis) Set(key string, value interface{}, expireDuration time.Duration) error {
	if r.pClient == nil {
		logger.Error.Fatalln("pClient is nil")
	}
	return r.pClient.Set(key, value, expireDuration).Err()
}

func (r *MainRedis) RemoveKey(key string, errWhenRedisNil bool) error {
	err := r.pClient.Del(key).Err()
	if err == redis.Nil && !errWhenRedisNil {
		return nil
	}
	return nil
}
