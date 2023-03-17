package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(ctx context.Context, config string) (*Redis, error) {
	// TODO
	return &Redis{}, nil
}

func (r *Redis) XAdd(ctx context.Context, channel uint32, data string) error {
	// TODO
	return nil
}

func (r *Redis) XGet(ctx context.Context, channel uint32) (<-chan string, error) {
	// TODO:
	return nil, nil
}

func (r *Redis) XTrim(ctx context.Context, channel uint32) error {
	// TODO
	return nil
}
