package event

import (
	"context"
	"time"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type producer struct {
	trim   int64
	expire time.Duration
	client *redis.Client
}

func NewProducer(ctx context.Context, config *core.Config) (Producer, error) {
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

	trim := config.Event.Trim
	if trim < core.StreamMaxLen {
		trim = core.StreamMaxLen
	}
	expire := time.Hour * time.Duration(config.Event.Expire)
	if expire < core.StreamExpire {
		expire = core.StreamExpire
	}

	log.Infof("producer trim %d exprie %0.f hour", trim, expire.Hours())
	return &producer{
		trim:   trim,
		expire: expire,
		client: client,
	}, nil
}

func (p *producer) Produce(ctx context.Context, stream string, data map[string]interface{}) (string, error) {
	args := &redis.XAddArgs{
		Stream:     stream,
		NoMkStream: false,
		ID:         "*",
		MaxLen:     p.trim,
		Values:     data,
	}

	pipe := p.client.Pipeline()
	defer pipe.Close()

	xAdd := pipe.XAdd(ctx, args)
	pipe.Expire(ctx, stream, p.expire)

	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Errorf("XAdd pipe Exec err %+v", err)
		return "", err
	}

	offset, err := xAdd.Result()
	if err != nil {
		log.Errorf("XAdd %+v err %+v", args, err)
		return "", err
	}
	log.Debugf("XAdd %s id %s", stream, offset)
	return offset, nil
}

func (p *producer) Close(ctx context.Context) error {
	return p.client.Close()
}
