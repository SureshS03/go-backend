package redis

import "time"

func SetCache(key string, value string, expiration time.Duration) error {
	return Client.Set(Ctx, key, value, expiration).Err()
}

func GetCache(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

func DeleteCache(key string) error {
	return Client.Del(Ctx, key).Err()
}

func CacheExists(key string) bool {
	count, err := Client.Exists(Ctx, key).Result()
	return err == nil && count > 0
}
