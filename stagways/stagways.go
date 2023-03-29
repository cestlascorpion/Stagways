package stagways

import (
	"context"

	"github.com/cestlascorpion/Stagways/core"
	log "github.com/sirupsen/logrus"
)

type Stagways struct {
	dao *event.Redis
}

func NewStagways(ctx context.Context, config *core.Config) (*Stagways, error) {
	log.Debugf("config %+v", config)

	dao, err := event.NewRedis(ctx, config)
	if err != nil {
		log.Errorf("new redis err %+v", err)
		return nil, err
	}

	return &Stagways{
		dao: dao,
	}, nil
}

func (s *Stagways) Publish(ctx context.Context, req *core.Message) error {
	// TODO:
	return nil
}

func (s *Stagways) Subscribe(ctx context.Context, streamId uint32) (<-chan *core.Message, error) {
	// TODO:
	return nil, nil
}

func (s *Stagways) Close(ctx context.Context) error {
	// TODO:
	return nil
}
