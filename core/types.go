package core

import "context"

type StreamType uint32

type Message struct {
	timestamp  int64
	streamId   uint32
	streamType StreamType
	taskId     string
	msgData    []byte
}

type Queue interface {
	Publish(ctx context.Context, req *Message) error
	Subscribe(ctx context.Context, streamId uint32) (<-chan *Message, error)
	Close(ctx context.Context) error
}
