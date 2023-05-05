package mq

import (
	"context"
)

// Producer produce single msg
type Producer interface {
	// Produce a single msg into the given stream, return offset
	// auto create stream if not exists
	Produce(ctx context.Context, stream string, data map[string]interface{}) (string, error)
	// Close producer
	Close(ctx context.Context) error
}

// Consumer SINGLE consumer mode: group name = consumer name
type Consumer interface {
	// Claim consumer group and bind stream, return err if stream not exists
	// always call Claim() before Consume()
	Claim(ctx context.Context, stream, group string) error
	// Check pending list and return size(should be zero)
	Check(ctx context.Context) (int, error)
	// Consume data from the stream(bind when Claim()) and latest offset
	// auto ack offset to make consumer more effective
	Consume(ctx context.Context) (map[string]interface{}, string, error)
	// Close consumer
	Close(ctx context.Context) error
}

// Manager manage a lot of consumer group
type Manager interface {
	// AddStream Allow repeated calls
	AddStream(ctx context.Context, stream string) error
	// DelStream Avoid repeated calls
	DelStream(ctx context.Context, stream string) error
	// Close all consumer
	Close(ctx context.Context) error
}
