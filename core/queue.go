package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type queue struct {
	rds *redis.Client
}

func NewQueueWithRedis(ctx context.Context, opts *redis.Options) (*queue, error) {
	rdb := redis.NewClient(opts)

	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Errorf("NewQueue rdb.Ping err %+v", err)
		return nil, err
	}

	return &queue{
		rds: rdb,
	}, nil
}

func (q *queue) Push(ctx context.Context, data *Meta) error {
	return nil
}

func (q *queue) Pop(ctx context.Context) (*Meta, error) {
	return nil, nil
}

func (q *queue) BatchPush(ctx context.Context, data []*Meta) error {
	return nil
}

func (q *queue) BatchPop(ctx context.Context) ([]*Meta, error) {
	return nil, nil
}

func (q *queue) Size(ctx context.Context) (int, error) {
	return 0, nil
}

func (q *queue) Clean(ctx context.Context) error {
	return nil
}

func (q *queue) Close() error {
	err := q.rds.Close()
	if err != nil {
		log.Errorf("Close rds.Close err %+v", err)
		return err
	}
	return nil
}
