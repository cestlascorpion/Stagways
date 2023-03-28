package core

import "context"

type StreamType uint32

type Message struct {
	StreamId   uint32     `json:"stream_id"`
	StreamType StreamType `json:"stream_type"`
	SequenceId string     `json:"sequence_id"`
	MsgData    []byte     `json:"msg_data"`
	Timestamp  int64      `json:"timestamp"`
}

type Queue interface {
	Publish(ctx context.Context, message *Message) error
	Subscribe(ctx context.Context, streamId []uint32) (<-chan *Message, error)
	Close(ctx context.Context) error
}
