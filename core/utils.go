package core

import (
	"fmt"
	"time"

	"github.com/cestlascorpion/Stagways/proto"
)

func GetStreamId(channel uint32) string {
	return fmt.Sprintf("stream_%d", channel)
}

func GetMessage(message *proto.Message) map[string]interface{} {
	result := make(map[string]interface{})
	result["task"] = message.PushTask
	result["type"] = message.PushType
	result["body"] = message.MsgBody
	result["req_ts"] = message.Timestamp
	result["in_ts"] = time.Now().UnixMilli()
	return result
}

type Message struct {
	Task  string `json:"task"`
	Type  uint32 `json:"type"`
	Body  []byte `json:"body"`
	ReqTs int64  `json:"req_ts"`
	InTs  int64  `json:"in_ts"`
	OutTs int64  `json:"out_ts"`
}

func ParseMessage(data map[string]interface{}) *Message {
	outTs := time.Now().UnixMilli()

	pushTask, ok := data["task"].(string)
	if !ok {
		pushTask = "unknown task"
	}
	pushType, ok := data["type"].(uint32)
	if !ok {
		pushType = 0
	}
	msgBody, ok := data["body"].([]byte)
	if !ok {
		// TODO:
	}
	reqTs, ok := data["req_ts"].(int64)
	if !ok {
		reqTs = 0
	}
	inTs, ok := data["in_ts"].(int64)
	if !ok {
		inTs = 0
	}

	return &Message{
		Task:  pushTask,
		Type:  pushType,
		Body:  msgBody,
		ReqTs: reqTs,
		InTs:  inTs,
		OutTs: outTs,
	}
}
