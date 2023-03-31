package agent

import (
	"context"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/cestlascorpion/Stagways/mq"
	"github.com/cestlascorpion/Stagways/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*proto.UnimplementedAgentServer
	config  *core.Config
	manager mq.Manager
}

func NewServer(ctx context.Context, conf *core.Config) (*Server, error) {
	mgr, err := mq.NewConsumerManager(ctx, conf, handler)
	if err != nil {
		log.Errorf("new consumer manager err %+v", err)
		return nil, err
	}

	return &Server{
		config:  conf,
		manager: mgr,
	}, nil
}

func (s *Server) AddChannel(ctx context.Context, in *proto.AddChannelReq) (*proto.AddChannelResp, error) {
	out := &proto.AddChannelResp{}

	stream := core.GetStreamId(in.Channel)
	err := s.manager.AddStream(ctx, stream)
	if err != nil {
		log.Errorf("manager add stream %s err %+v", stream, err)
		return out, err
	}

	return out, nil
}

func (s *Server) DelChannel(ctx context.Context, in *proto.DelChannelReq) (*proto.DelChannelResp, error) {
	out := &proto.DelChannelResp{}

	stream := core.GetStreamId(in.Channel)
	err := s.manager.DelStream(ctx, stream)
	if err != nil {
		log.Errorf("manager del stream %s err %+v", stream, err)
		return out, err
	}

	return out, nil
}

func (s *Server) Close(ctx context.Context) error {
	return s.manager.Close(ctx)
}

func handler(ctx context.Context, message map[string]interface{}, offset string) {
	msg := core.ParseMessage(message)
	log.Infof("offset %s task %s type %d body %v req %d in %d out %d", offset, msg.Task, msg.Type, msg.Body, msg.ReqTs, msg.InTs, msg.OutTs)
}
