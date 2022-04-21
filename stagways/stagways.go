package stagways

import (
	"context"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Stagways struct {
	queue core.Queue
	// TODO:
}

type Options struct {
	rdbOpts *redis.Options
	// TODO:
}

func NewStagways(ctx context.Context) (*Stagways, error) {
	opts := &Options{
		// TODO:
	}

	q, err := core.NewQueueWithRedis(ctx, opts.rdbOpts)
	if err != nil {
		log.Errorf("NewStagways core.NewQueueWithRedis err %+v", err)
		return nil, err
	}

	return &Stagways{
		queue: q,
	}, nil
}
