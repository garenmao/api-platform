package cache

import (
	"api-platform/global"
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"time"
)

type RedisStore struct {
	Prefix     string
	Context    context.Context
	Expiration time.Duration
}

var _ base64Captcha.Store = &RedisStore{}

var DefaultRedisStore = &RedisStore{
	Expiration: time.Minute * 15,
	Prefix:     "CAPTCHA_",
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	key := rs.Prefix + id
	err := global.Redis.Set(rs.Context, key, value, rs.Expiration).Err()
	if err != nil {
		fmt.Errorf("RedisStoreSetError:key:%s,value:%s", key, value)
	}
	return nil
}

func (rs *RedisStore) Get(id string, clear bool) string {
	key := rs.Prefix + id
	result, err := global.Redis.Get(rs.Context, key).Result()
	if err != nil {
		fmt.Errorf("get redis value error, %s", err)
		return ""
	}
	if clear {
		err := global.Redis.Del(rs.Context, key).Err()
		if err != nil {
			fmt.Errorf("delete redis value error, %s", err)
			return ""
		}
	}
	return result
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	value := rs.Get(id, clear)
	return value == answer
}
