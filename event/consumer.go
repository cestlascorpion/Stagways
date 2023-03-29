package event

import (
	"context"
	"errors"
	"strings"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type consumer struct {
	name   string
	stream string
	group  string
	client *redis.Client
}

func NewConsumer(ctx context.Context, config *core.Config) (Consumer, error) {
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

	return &consumer{
		client: client,
	}, nil
}

func (c *consumer) Claim(ctx context.Context, stream, group string) error {
	_, err := c.client.XGroupCreate(ctx, stream, group, "0").Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Errorf("XGroupCreate %s %s err %+v", stream, group, err)
		return err
	}

	c.name = uuid.New().String()
	c.stream = stream
	c.group = group
	log.Debugf("XGroupCreate stream %s group %s name %s", c.stream, c.group, c.name)
	return nil
}

func (c *consumer) Consume(ctx context.Context) (map[string]interface{}, string, error) {
	if len(c.group) == 0 || len(c.stream) == 0 {
		log.Errorf("consumer requires claim")
		return nil, "", errors.New("consumer requires claim")
	}

	result, err := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    c.group,
		Consumer: c.name,
		Streams:  []string{c.stream, ">"}, // TODO: use "0" when recovery from panic
		Count:    1,
		Block:    core.StreamBlock,
		NoAck:    false,
	}).Result()

	if err != nil {
		if err == redis.Nil {
			log.Debugf("XReadGroup %s return nothing", c.stream)
			return nil, "", nil
		}
		log.Errorf("XReadGroup %s err %+v", c.stream, err)
		return nil, "", err
	}

	if len(result) == 0 {
		log.Debugf("XReadGroup %s return nothing", c.stream)
		return nil, "", nil
	}

	for i := range result {
		if result[i].Stream != c.stream {
			log.Warnf("XReadGroup %s != %s", result[i].Stream, c.stream)
			continue
		}

		if len(result[i].Messages) == 0 {
			log.Debugf("XReadGroup %s return nothing", c.stream)
			return nil, "", nil
		}

		log.Debugf("XReadGroup %s msg %v offset %s",
			c.stream, result[i].Messages[0].Values, result[i].Messages[0].ID)
		return result[i].Messages[0].Values, result[i].Messages[0].ID, nil
	}

	log.Debugf("XReadGroup %s return nothing", c.stream)
	return nil, "", nil
}

func (c *consumer) Ack(ctx context.Context, offset string) error {
	_, err := c.client.XAck(ctx, c.stream, c.group, offset).Result()
	if err != nil {
		log.Errorf("XAck %s %s %s err %+v", c.stream, c.group, offset, err)
		return err
	}

	log.Debugf("XAck %s %s %s", c.stream, c.group, offset)
	return nil
}

func (c *consumer) Close(ctx context.Context) error {
	return c.client.Close()
}
