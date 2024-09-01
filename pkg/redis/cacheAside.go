package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/ongyoo/roomkub-api/internal/metric"
)

func GetCacheAside[T any](ctx context.Context, name string, rdb redis.UniversalClient, key string, ttl time.Duration, fetchData func() (*T, error)) (*T, error) {
	jsonStr, err := rdb.Get(ctx, key).Result()

	var result T
	// cache hit, so return cached result
	if err == nil {
		metric.IncCacheHit(name)
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err == nil {
			return &result, nil
		}
		// If there is any error when unmarshalling maybe due to corrupt JSON, we will fetch the data from the database
	}

	// regardless of cache miss or redis error, fetch the data from the database
	metric.IncCacheMiss(name)
	value, err := fetchData()
	if err != nil {
		return nil, err
	}

	// Asynchonously store the data in the cache for future use
	// even if it fails, still return successful data
	go func() {
		valueJson, err := json.Marshal(value)
		if err != nil {

		}
		err = rdb.Set(context.Background(), key, string(valueJson), ttl).Err()
		if err != nil {

		}
	}()

	return value, nil
}
