package event

import (
	"context"
)

type Producer interface {
	// Produce write data into the given steam, return offset
	Produce(ctx context.Context, stream string, data map[string]interface{}) (string, error)
	// Close producer
	Close(ctx context.Context) error
}

type Consumer interface {
	// Claim consumer group and bind stream, return err if stream not exists
	Claim(ctx context.Context, stream, group string) error
	// Consume read data from the stream(bind when Create()) and latest offset
	Consume(ctx context.Context) (map[string]interface{}, string, error)
	// Ack the offset
	Ack(ctx context.Context, offset string) error
	// Close consumer
	Close(ctx context.Context) error
}
