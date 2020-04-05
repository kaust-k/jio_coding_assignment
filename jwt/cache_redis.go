package jwt

import (
	"fmt"
	"jwt_server/cache"
	"time"
)

type redisCacheProvider struct {
}

func (cp *redisCacheProvider) save(e *entry) error {
	client := cache.GetRedisClient()
	key := cp.getKey(e)
	now := time.Now().UTC()
	dur := e.Expires.Sub(now)
	return client.Set(key, "", dur).Err()
}

func (cp *redisCacheProvider) delete(e *entry) error {
	client := cache.GetRedisClient()
	key := cp.getKey(e)
	return client.Del(key).Err()
}

func (cp *redisCacheProvider) isValid(e *entry) (bool, error) {
	client := cache.GetRedisClient()
	key := cp.getKey(e)
	_, err := client.Get(key).Result()
	return err != nil, err
}

func (cp *redisCacheProvider) getKey(e *entry) string {
	return fmt.Sprintf("%s:%s:jwt", e.UserID, e.AuthID)
}
