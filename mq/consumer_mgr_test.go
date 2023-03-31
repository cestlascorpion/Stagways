package mq

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConsumerManager_AddStream(t *testing.T) {
	if testManager == nil {
		return
	}

	err := testManager.AddStream(context.Background(), "s1")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	err = testManager.AddStream(context.Background(), "s2")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	err = testManager.AddStream(context.Background(), "s3")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	err = testManager.AddStream(context.Background(), "s1")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	_ = testManager.Close(context.Background())
}

func TestConsumerManager_DelStream(t *testing.T) {
	if testManager == nil {
		return
	}

	err := testManager.AddStream(context.Background(), "s1")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	err = testManager.DelStream(context.Background(), "s1")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	err = testManager.DelStream(context.Background(), "s1")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	_ = testManager.Close(context.Background())
}

func TestConsumerManager_Close(t *testing.T) {
	if testManager == nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			err := testManager.AddStream(context.Background(), fmt.Sprintf("s-%d", i))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("add s -", i)
			}
		}
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second)

		for i := 0; i < 10; i++ {
			err := testManager.DelStream(context.Background(), fmt.Sprintf("s-%d", i))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("del s -", i)
			}
		}
	}()

	wg.Wait()
}
