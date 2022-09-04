package repository

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisAdapter() (*DbAdapter, error) {
	otp, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(otp)

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return &DbAdapter{
		redisClient: redisClient,
	}, nil
}

// Keys exspire after 180 seconds
func (a *DbAdapter) Set(ctx context.Context, key string, session interface{}) error {
	err := a.redisClient.Set(ctx, key, session, 180*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) Get(ctx context.Context, key string) (string, error) {
	val, err := a.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		//Key doesnt exist
		return "", err
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (a *DbAdapter) Delete(ctx context.Context, key string) error {
	err := a.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (a *DbAdapter) Update(ctx context.Context, key string) error {
	err := a.redisClient.Expire(ctx, key, 180*time.Second).Err()
	if err == redis.Nil {
		//Key doesnt exist
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
