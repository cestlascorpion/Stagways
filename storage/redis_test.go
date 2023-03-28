package storage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cestlascorpion/Stagways/core"
)

func TestRedis_XAdd(t *testing.T) {
	conf, err := core.NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dao, err := NewRedis(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	result, err := dao.XAdd(context.Background(), "12345678", map[string]interface{}{"ts": time.Now().String()},
		core.StreamMaxLen, core.StreamTTL)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(result)

	_ = dao.Close(context.Background())
}

func TestRedis_XRevRange(t *testing.T) {
	conf, err := core.NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dao, err := NewRedis(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	result, id, err := dao.XRevRange(context.Background(), "12345678", core.StreamLimit)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(id)
	fmt.Println(result)
}

func BenchmarkRedis_XAdd(b *testing.B) {
	conf, err := core.NewConfig(context.Background(), "/home/hans/Workspace/Stagways/config.json")
	if err != nil {
		fmt.Println(err)
		b.FailNow()
	}

	dao, err := NewRedis(context.Background(), conf)
	if err != nil {
		fmt.Println(err)
		b.FailNow()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := dao.XAdd(context.Background(), "12345678", map[string]interface{}{"key": "value"},
			core.StreamMaxLen, core.StreamTTL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}

	b.StopTimer()
	_ = dao.Close(context.Background())
}
