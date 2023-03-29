package core

import (
	"context"

	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server struct {
		Name   string `json:"name"`
		Listen string `json:"listen"`
	} `json:"server"`
	Redis struct {
		Addr     string `json:"addr"`
		Protocol string `json:"protocol"`
		Database int    `json:"database"`
		PoolSize int    `json:"pool_size"`
	} `json:"redis"`
	Event struct {
		Trim   int64 `json:"trim"`
		Expire int64 `json:"expire"`
	} `json:"event"`
}

func NewConfig(ctx context.Context, path string) (*Config, error) {
	conf := &Config{}
	err := configor.Load(conf, path)
	if err != nil {
		log.Errorf("load config err %+v", err)
		return nil, err
	}
	return conf, nil
}
