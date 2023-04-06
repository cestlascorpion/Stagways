package proxy

import (
	"context"
	"errors"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/cestlascorpion/Stagways/mq"
	"github.com/cestlascorpion/Stagways/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*proto.UnimplementedProxyServer
	producer mq.Producer
}

func NewServer(ctx context.Context, conf *core.Config) (*Server, error) {
	p, err := mq.NewProducer(ctx, conf)
	if err != nil {
		log.Errorf("new producer err %+v", err)
		return nil, err
	}

	return &Server{
		producer: p,
	}, nil
}

func (s *Server) PushMessage(ctx context.Context, in *proto.PushMessageReq) (*proto.PushMessageResp, error) {
	out := &proto.PushMessageResp{}

	if in.Message == nil || len(in.Message.PushTask) == 0 || len(in.Message.MsgBody) == 0 {
		log.Errorf("invalid push message")
		return out, errors.New("invalid parameter")
	}

	stream := core.GetStreamId(in.Channel)
	message, err := core.GetMessage(in.Message)
	if err != nil {
		log.Errorf("get message err %+v", err)
		return out, err
	}
	receipt, err := s.producer.Produce(ctx, stream, message)
	if err != nil {
		log.Errorf("produce %s msg %v err %+v", stream, message, err)
		return out, err
	}

	out.Receipt = receipt
	return out, nil
}

func (s *Server) Close(ctx context.Context) error {
	return s.producer.Close(ctx)
}
