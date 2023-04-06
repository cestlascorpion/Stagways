package mq

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type consumer struct {
	expire time.Duration
	block  time.Duration
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

	expire := time.Second * time.Duration(config.Queue.ExpireSec)
	if expire < core.StreamExpire {
		expire = core.StreamExpire
	}
	block := time.Millisecond * time.Duration(config.Queue.BlockMs)
	if expire < core.StreamBlock {
		expire = core.StreamBlock
	}

	log.Infof("comsumer block %s exprie %s", block.String(), expire.String())
	return &consumer{
		expire: expire,
		block:  block,
		client: client,
	}, nil
}

func (c *consumer) Claim(ctx context.Context, stream, group string) error {
	pipe := c.client.Pipeline()
	defer pipe.Close()

	statusCmd := pipe.XGroupCreateMkStream(ctx, stream, group, "0")
	pipe.Expire(ctx, stream, c.expire)

	_, err := pipe.Exec(ctx)
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Errorf("Claim pipe Exec err %+v", err)
		return err
	}

	_, err = statusCmd.Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP") {
		log.Errorf("XGroupCreate stream %s group %s err %+v", stream, group, err)
		return err
	}

	c.stream = stream
	c.group = group
	log.Debugf("XGroupCreate stream %s group %s", c.stream, c.group)
	return nil
}

func (c *consumer) Check(ctx context.Context) (int, error) {
	if len(c.group) == 0 || len(c.stream) == 0 {
		log.Errorf("consumer requires claim")
		return 0, errors.New("consumer requires claim")
	}

	args := &redis.XPendingExtArgs{
		Stream:   c.stream,
		Group:    c.group,
		Start:    "-",
		End:      "+",
		Count:    1024,
		Consumer: c.group,
	}

	result, err := c.client.XPendingExt(ctx, args).Result()
	if err != nil {
		log.Errorf("XPending stream %s group %s err %+v", c.stream, c.group, err)
		return 0, err
	}

	if len(result) == 0 {
		log.Debugf("XPending stream %s no pending msg", c.stream)
		return 0, nil
	}

	idList := make([]string, 0, len(result))
	for i := range result {
		log.Warnf("XPending %s %s %s %d", result[i].Consumer, result[i].ID, result[i].Idle, result[i].RetryCount)
		idList = append(idList, result[i].ID)
	}

	count, err := c.client.XAck(ctx, c.stream, c.group, idList...).Result()
	if err != nil {
		log.Warnf("XAck stream %s group %s %v err %+v", c.stream, c.group, idList, err)
	} else {
		log.Infof("XAck stream %s group %s %d %d", c.stream, c.group, len(result), count)
	}

	return int(count), nil
}

func (c *consumer) Consume(ctx context.Context) (map[string]interface{}, string, error) {
	if len(c.group) == 0 || len(c.stream) == 0 {
		log.Errorf("consumer requires claim")
		return nil, "", errors.New("consumer requires claim")
	}

	result, err := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    c.group,
		Consumer: c.group,
		Streams:  []string{c.stream, ">"},
		Count:    1,
		Block:    c.block,
		NoAck:    true,
	}).Result()

	if err != nil {
		if err == redis.Nil {
			log.Debugf("XReadGroup stream %s return nothing", c.stream)
			return nil, "", nil
		}
		log.Errorf("XReadGroup stream %s err %+v", c.stream, err)
		return nil, "", err
	}

	if len(result) != 1 {
		log.Debugf("XReadGroup stream %s return nothing", c.stream)
		return nil, "", nil
	}

	if len(result[0].Messages) != 1 {
		log.Debugf("XReadGroup stream %s return nothing", c.stream)
		return nil, "", nil
	}

	log.Debugf("XReadGroup stream %s offset %s", c.stream, result[0].Messages[0].ID)
	return result[0].Messages[0].Values, result[0].Messages[0].ID, nil
}

func (c *consumer) Close(ctx context.Context) error {
	return c.client.Close()
}
