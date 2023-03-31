package mq

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestProducer_Produce(t *testing.T) {
	if testProducer == nil {
		return
	}

	for i := int64(0); i < 10; i++ {
		offset, err := testProducer.Produce(context.Background(), testSteam, map[string]interface{}{
			"id": strconv.FormatInt(i, 10),
		})
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		fmt.Println(i, offset)
	}

	_ = testProducer.Close(context.Background())
}
