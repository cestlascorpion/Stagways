package mq

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/cestlascorpion/Stagways/core"
	log "github.com/sirupsen/logrus"
)

var (
	testConsumer Consumer
	testProducer Producer
	testManager  Manager
	testGroup    string
	testSteam    string
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})

	conf, err := core.NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := NewConsumer(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err := NewProducer(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := NewConsumerManager(context.Background(), conf, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	testConsumer = c
	testProducer = p
	testManager = m
	testGroup = "GROUP"
	testSteam = "STREAM"
}

func Test_MQ(t *testing.T) {
	if testConsumer == nil || testProducer == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	x, cancel := context.WithTimeout(context.Background(), time.Second*10)

	go func(ctx context.Context) {
		defer wg.Done()

		err := testConsumer.Claim(ctx, testSteam, testGroup)
		if err != nil {
			fmt.Println(err)
			return
		}

		count, err := testConsumer.Check(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("clean msg", count)

		for {
			msg, offset, err := testConsumer.Consume(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("consume:", offset, msg)
		}
	}(x)

	go func(ctx context.Context) {
		defer wg.Done()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				offset, err := testProducer.Produce(ctx, testSteam, map[string]interface{}{
					"id": time.Now().Unix(),
				})
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("produce", offset)
			}
		}
	}(x)

	time.Sleep(time.Second * 12)

	cancel()
	wg.Wait()
}
