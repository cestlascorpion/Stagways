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
	noAck  bool
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

func (c *consumer) Claim(ctx context.Context, stream, group string, noAck bool) error {
	_, err := c.client.XGroupCreate(ctx, stream, group, "0").Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Errorf("XGroupCreate %s %s err %+v", stream, group, err)
		return err
	}
	c.name = uuid.New().String()
	c.stream = stream
	c.group = group
	c.noAck = noAck
	log.Debugf("XGroupCreate stream %s group %s name %s noAck %v", c.stream, c.group, c.name, noAck)
	return nil
}

func (c *consumer) Consume(ctx context.Context) ([]map[string]interface{}, string, error) {
	if len(c.group) == 0 || len(c.stream) == 0 {
		log.Errorf("consumer requires claim")
		return nil, "", errors.New("consumer requires claim")
	}

	result, err := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    c.group,
		Consumer: c.name,
		Streams:  []string{c.stream, ">"},
		Count:    core.StreamLimit,
		Block:    core.StreamBlock,
		NoAck:    c.noAck,
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

	dataList := make([]map[string]interface{}, 0, len(result))
	offset := ""
	for i := range result {
		if result[i].Stream == c.stream {
			for j := range result[i].Messages {
				dataList = append(dataList, result[i].Messages[j].Values)
				offset = result[i].Messages[j].ID
			}
		}
	}
	log.Debugf("XReadGroup %s offset %s", c.stream, offset)
	return dataList, offset, nil
}

func (c *consumer) Ack(ctx context.Context, offset string) error {
	return nil
}

func (c *consumer) Close(ctx context.Context) error {
	return c.client.Close()
}
