package infra

import (
	"context"
	"sync"
	"time"

	"github.com/patrickishaf/job_scheduler/config"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	mu     *sync.Mutex
}

func InitRedisCache(cfg *config.RedisConfig) *RedisCache {
	var mu sync.Mutex
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.GetAddr(),
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
		mu: &mu,
	}
}

func (this *RedisCache) Dequeue(ctx context.Context, key string) (string, error) {
	b, err := this.client.LPop(ctx, key).Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (this *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return this.client.Get(ctx, key).Result()
}

func (this *RedisCache) Queue(ctx context.Context, key string, value any) error {
	return this.client.RPush(ctx, key, value).Err()
}

func (this *RedisCache) Set(ctx context.Context, key string, value string, ttlHours int) error {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.client.Set(ctx, key, value, time.Duration(ttlHours*int(time.Hour))).Err()
}
