package main

import (
	"context"
	"flag"
	"net"

	"github.com/cestlascorpion/Stagways/core"
	"github.com/cestlascorpion/Stagways/proto"
	"github.com/cestlascorpion/Stagways/service/agent"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	cfg      string
	logLevel string
)

func init() {
	flag.StringVar(&cfg, "config", "config.json", "config path")
	flag.StringVar(&logLevel, "logLevel", "debug", "log level")
}

func main() {
	flag.Parse()

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("parse level err %+v", err)
		return
	}
	log.SetLevel(level)

	conf := &core.Config{}
	err = configor.Load(conf, cfg)
	if err != nil {
		log.Fatalf("config failed err %+v", err)
		return
	}

	lis, err := net.Listen("tcp", conf.Server.Agent)
	if err != nil {
		log.Fatalf("listen failed err %+v", err)
		return
	}

	ctx := context.Background()
	svr, err := agent.NewServer(ctx, conf)
	if err != nil {
		log.Fatalf("new server failed err %+v", err)
		return
	}
	defer svr.Close(ctx)

	s := grpc.NewServer()
	proto.RegisterAgentServer(s, svr)
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("serve failed err %+v", err)
		return
	}
}
