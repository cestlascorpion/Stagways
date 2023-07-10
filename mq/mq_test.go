package mq

import (
	"context"
	"fmt"
	"strconv"
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
			if msg != nil {
				fmt.Println("consume:", offset, msg)
			}
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

func Test_MQ_2(t *testing.T) {
	conf, err := core.NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	parallelP := 300
	parallelC := 50
	waitSec := 5
	consumerList := make([]Consumer, 0)
	for i := 0; i < parallelC; i++ {

		c, err := NewConsumer(context.Background(), conf)
		if err != nil {
			fmt.Println(err)
			return
		}

		consumerList = append(consumerList, c)
	}

	p, err := NewProducer(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(parallelP + parallelC)

	x, cancel := context.WithTimeout(context.Background(), time.Second*30)

	for i := 0; i < parallelC; i++ {
		go func(ctx context.Context, idx int) {
			defer wg.Done()

			err = consumerList[idx].Claim(ctx, testSteam, fmt.Sprintf("consumer-%d", idx))
			if err != nil {
				fmt.Println(err)
				return
			}

			_, err = consumerList[idx].Check(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}

			for {
				msg, _, e := consumerList[idx].Consume(ctx)
				if e != nil {
					fmt.Println(err)
					return
				}
				if msg != nil {
					data, ok := msg["ts"]
					if !ok {
						fmt.Println("no ts found")
						continue
					}
					val, ok := data.(string)
					if !ok {
						fmt.Println("unknown type")
						continue
					}
					ts, e := strconv.ParseInt(val, 10, 64)
					if e != nil {
						fmt.Println(e)
						continue
					}
					fmt.Println(float64((time.Now().UnixNano()-ts)/1e6), " ms")
				}
			}
		}(x, i)
	}

	time.Sleep(time.Second * time.Duration(waitSec))

	for i := 0; i < parallelP; i++ {
		go func(ctx context.Context, idx int) {
			defer wg.Done()

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					_, _ = p.Produce(ctx, testSteam, map[string]interface{}{
						"ts": time.Now().UnixNano(),
					})
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}(x, i)
	}
	time.Sleep(time.Second * 60)

	cancel()
	wg.Wait()
}
