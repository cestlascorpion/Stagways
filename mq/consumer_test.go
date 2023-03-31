package mq

import (
	"context"
	"fmt"
	"testing"
)

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

func TestConsumer_Check(t *testing.T) {
	if testConsumer == nil {
		return
	}

	err := testConsumer.Claim(context.Background(), testSteam, testGroup)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	count, err := testConsumer.Check(context.Background())
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(count)
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

	data, offset, err := testConsumer.Consume(context.Background())
	if err != nil || len(offset) == 0 {
		fmt.Println(err)
		t.FailNow()
	}
	if data != nil {
		fmt.Println(offset)
		fmt.Println(data)
	}
}
