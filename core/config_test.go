package core

import (
	"context"
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	conf, err := NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(conf.Server, conf.Redis, conf.Queue)
}
