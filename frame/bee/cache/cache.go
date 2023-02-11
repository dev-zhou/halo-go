package cache

import (
	"HaloAdmin/conf"
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"time"
)

var Redis cache.Cache

func init() {
	var err error
	config := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%v","password":"%s"}`,
		"redis_Tokey", conf.RedisAddr, conf.RedisDBNo, conf.RedisPass)
	Redis, err = cache.NewCache("redis", config)

	if err != nil || Redis == nil {
		fmt.Println("init cache failed", err)
	}
}

func Set(key string, value interface{}) error {
	timeout := 24 * time.Hour
	err := Redis.Put(context.Background(), key, value, timeout)
	if err != nil {
		fmt.Printf("halo/frame/bee/cache -- Set Error:%#v \n", err)
	}
	return err
}

func Get(key string) (get interface{}, err error) {
	get, err = Redis.Get(context.Background(), key)
	if err != nil {
		fmt.Printf("halo/frame/bee/cache -- Get Error:%#v \n", err)
		return
	}
	return get, err
}

func Del(key string) error {
	err := Redis.Delete(context.Background(), key)
	if err != nil {
		fmt.Printf("halo/frame/bee/cache -- Get Error:%#v \n", err)
	}
	return err
}

func Has(key string) (bool, error) {
	exist, err := Redis.IsExist(context.Background(), key)
	return exist, err
}
