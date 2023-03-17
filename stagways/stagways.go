package stagways

import (
	"context"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/cestlascorpion/Stagways/storage"
	log "github.com/sirupsen/logrus"
)

type Stagways struct {
	dao *storage.Redis
}

func NewStagways(ctx context.Context, config string) (*Stagways, error) {
	log.Debugf("config %s", config)

	// TODO:

	return &Stagways{}, nil
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
