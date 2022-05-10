package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/wujunyi792/where-it-go/config"
	"github.com/wujunyi792/where-it-go/internal/logger"
	"time"
)

type systemCache struct {
	useRedis bool
	client   interface{}
}

var autoCache *systemCache

func init() {
	autoInit()
}

func autoInit() {
	if config.GetConfig().REDIS.Use {
		initRedis()
	} else {
		initGoCache()
	}
}

func initRedis() {
	conf := &config.GetConfig().REDIS.Config
	var client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.IP, conf.PORT),
		Password: conf.PASSWORD,
		DB:       conf.DB,
	})
	autoCache = &systemCache{
		useRedis: true,
		client:   client,
	}
	_, err := client.Ping().Result()
	if err != nil {
		logger.Error.Fatalln(err)
		return
	}
	logger.Info.Printf("redis init finish")
}

func initGoCache() {
	autoCache = &systemCache{
		useRedis: false,
		client:   cache.New(cache.NoExpiration, 10*time.Minute),
	}
	logger.Info.Printf("local cache init finish")
}

func GetCache() *systemCache {
	if autoCache == nil {
		autoInit()
	}
	return autoCache
}

func (c *systemCache) Set(key string, value interface{}, expireDuration time.Duration) error {
	if c.client == nil {
		logger.Error.Println("cache client fatal error")
		return errors.New("cant set value")
	}
	if c.useRedis {
		return c.client.(*redis.Client).Set(key, value, expireDuration).Err()
	}
	c.client.(*cache.Cache).Set(key, value, expireDuration)
	return nil
}

func (c *systemCache) RemoveKey(key string, errWhenClientNil bool) error {
	if c.useRedis {
		err := c.client.(*redis.Client).Del(key).Err()
		if err == redis.Nil && !errWhenClientNil {
			return nil
		} else {
			return errors.New("redis client offline")
		}
	}
	c.client.(*cache.Cache).Delete(key)
	return nil
}

// Get empty result throw error
func (c *systemCache) Get(key string) (string, error) {
	if c.client == nil {
		logger.Error.Println("cache client fatal error")
		return "", errors.New("cache server error")
	}
	if c.useRedis {
		return c.client.(*redis.Client).Get(key).Result()
	}
	if x, found := c.client.(*cache.Cache).Get(key); found {
		return x.(string), nil
	} else {
		return "", errors.New("record not found")
	}
}
