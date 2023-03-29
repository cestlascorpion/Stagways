package event

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cestlascorpion/Stagways/core"
	log "github.com/sirupsen/logrus"
)

var (
	testProducer Producer
	testSteam    string
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	conf, err := core.NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err := NewProducer(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	testProducer = p
	testSteam = "STREAM"
}

func TestProducer_Produce(t *testing.T) {
	if testProducer == nil {
		return
	}

	offset, err := testProducer.Produce(context.Background(), testSteam, map[string]interface{}{"ts": time.Now().Unix()})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(offset)

	_ = testProducer.Close(context.Background())
}
