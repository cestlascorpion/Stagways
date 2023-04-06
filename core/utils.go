package core

import (
	"errors"
	"fmt"

	pb "github.com/cestlascorpion/Stagways/proto"
	"google.golang.org/protobuf/proto"
)

func GetStreamId(channel uint32) string {
	return fmt.Sprintf("stream_%d", channel)
}

func GetMessage(message *pb.Message) (map[string]interface{}, error) {
	bs, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"msg": string(bs),
	}, nil
}

type Message struct {
	Task  string `json:"task"`
	Type  uint32 `json:"type"`
	Body  string `json:"body"`
	ReqTs int64  `json:"req_ts"`
	InTs  int64  `json:"in_ts"`
	OutTs int64  `json:"out_ts"`
}

func ParseMessage(data map[string]interface{}) (*pb.Message, error) {
	msg, ok := data["msg"]
	if !ok {
		return nil, errors.New("msg not found")
	}
	bs, ok := msg.(string)
	if !ok {
		return nil, errors.New("reflect failed")
	}
	message := &pb.Message{}
	err := proto.Unmarshal([]byte(bs), message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
