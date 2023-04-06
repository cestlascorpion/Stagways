package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cestlascorpion/Stagways/proto"
	"google.golang.org/grpc"
)

func NewAgent() (proto.AgentClient, error) {
	conn, err := grpc.Dial("localhost:8090", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	agent := proto.NewAgentClient(conn)
	return agent, nil
}

func NewProxy() (proto.ProxyClient, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	proxy := proto.NewProxyClient(conn)
	return proxy, nil
}

func main() {
	agent, err := NewAgent()
	if err != nil {
		fmt.Println(err)
		return
	}

	proxy, err := NewProxy()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = agent.AddChannel(context.Background(), &proto.AddChannelReq{
		Channel: 12345678,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		resp, err := proxy.PushMessage(context.Background(), &proto.PushMessageReq{
			Message: &proto.Message{
				PushTask:  fmt.Sprintf("task-%s-%d", time.Now().Format("15:04:05"), i),
				PushType:  1,
				MsgBody:   "hello,world!",
				Timestamp: time.Now().UnixMilli(),
			},
			Channel: 12345678,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.Receipt)
		time.Sleep(time.Millisecond * 10)
	}

	time.Sleep(time.Second * 5)

	_, err = agent.DelChannel(context.Background(), &proto.DelChannelReq{
		Channel: 12345678,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("test done...")
}
