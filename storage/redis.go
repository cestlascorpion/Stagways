package storage

import (
	"context"
	"time"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(ctx context.Context, config *core.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Network:  config.Redis.Protocol,
		DB:       config.Redis.Database,
		PoolSize: config.Redis.PoolSize,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Errorf("redis ping err %+v", err)
		return nil, err
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) XAdd(ctx context.Context, stream string, data map[string]interface{}, trim int64, ttl time.Duration) (string, error) {
	if len(stream) == 0 {
		log.Errorf("XAdd empty stream")
		return "", core.ErrEmptyStream
	}
	if len(data) == 0 {
		log.Errorf("XAdd empty data")
		return "", core.ErrEmptyData
	}
	if trim < core.StreamMaxLen {
		trim = core.StreamMaxLen
	}
	if ttl < core.StreamTTL {
		ttl = core.StreamTTL
	}

	args := &redis.XAddArgs{
		Stream:     stream,
		NoMkStream: false,
		ID:         "*",
		MaxLen:     trim,
		Values:     data,
	}

	pipe := r.client.Pipeline()
	defer pipe.Close()

	xAdd := pipe.XAdd(ctx, args)
	pipe.Expire(ctx, stream, ttl)

	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Errorf("XAdd pipe Exec err %+v", err)
		return "", err
	}

	id, err := xAdd.Result()
	if err != nil {
		log.Errorf("XAdd %+v err %+v", args, err)
		return "", err
	}
	log.Debugf("XAdd %s id %s", stream, id)
	return id, nil
}

func (r *Redis) XRevRange(ctx context.Context, stream string, limit int64) ([]map[string]interface{}, string, error) {
	if len(stream) == 0 {
		log.Errorf("XRevRange empty stream")
		return nil, "", core.ErrEmptyStream
	}
	if limit < core.StreamLimit {
		limit = core.StreamLimit
	}

	msgList, err := r.client.XRevRangeN(ctx, stream, "+", "-", limit).Result()
	if err != nil {
		log.Errorf("XRevRange %s err %+v", stream, err)
		return nil, "", err
	}

	if len(msgList) == 0 {
		log.Debugf("XRevRange %s no data", stream)
		return nil, "", nil
	}

	result := make([]map[string]interface{}, 0, len(msgList))
	id := msgList[0].ID
	for i := range msgList {
		result = append(result, msgList[i].Values)
	}
	log.Debugf("XRevRange %s id %s %d", stream, id, len(result))
	return result, id, nil
}

func (r *Redis) XAck(ctx context.Context, stream, id string) error {
	return nil
}

func (r *Redis) Close(ctx context.Context) error {
	return r.client.Close()
}
