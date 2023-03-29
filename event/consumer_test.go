package event

import (
	"context"
	"fmt"
	"testing"

	"github.com/cestlascorpion/Stagways/core"
	log "github.com/sirupsen/logrus"
)

var (
	testConsumer Consumer
	testGroup    string
)

func init() {
	log.SetLevel(log.DebugLevel)
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

	testConsumer = c
	testGroup = "GROUP"
}

func TestConsumer_Claim(t *testing.T) {
	if testConsumer == nil {
		return
	}

	err := testConsumer.Claim(context.Background(), testSteam, testGroup)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestConsumer_Consume(t *testing.T) {
	if testConsumer == nil {
		return
	}

	err := testConsumer.Claim(context.Background(), testSteam, testGroup)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dataList, offset, err := testConsumer.Consume(context.Background())
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(offset)
	fmt.Println(dataList)
}
