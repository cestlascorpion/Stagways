package core

import "context"

// Meta meta data of queue
type Meta struct {
	Topic string      `json:"topic"`
	Shard int         `json:"shard"`
	Data  interface{} `json:"data"`
}

// Queue just a simple queue(for a cpp coder)
type Queue interface {
	Push(ctx context.Context, data *Meta) error
	Pop(ctx context.Context) (*Meta, error)
	BatchPush(ctx context.Context, data []*Meta) error
	BatchPop(ctx context.Context) ([]*Meta, error)
	Size(ctx context.Context) (int, error)
	Clean(ctx context.Context) error
}
