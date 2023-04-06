package core

import (
	"context"
	"errors"

	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Server *server `json:"server,omitempty"`
	Redis  *redis  `json:"redis,omitempty"`
	Queue  *queue  `json:"queue,omitempty"`
}

type server struct {
	Proxy string `json:"proxy"` // proxy listen addr -> producer
	Agent string `json:"agent"` // agent listen addr -> consumer
}

type redis struct {
	Addr     string `json:"addr"`
	Protocol string `json:"protocol"`
	Database int    `json:"database"`
	PoolSize int    `json:"pool_size"`
}

type queue struct {
	MaxLen    int64  `json:"max_len"`
	ExpireSec int64  `json:"expire_sec"`
	BlockMs   int64  `json:"block_ms"`
	GroupName string `json:"group_name"`
}

func NewConfig(ctx context.Context, path string) (*Config, error) {
	conf := &Config{}
	err := configor.Load(conf, path)
	if err != nil {
		log.Errorf("load config err %+v", err)
		return nil, err
	}
	err = conf.check()
	if err != nil {
		log.Errorf("check config err %+v", err)
		return nil, err
	}
	return conf, nil
}

func (c *Config) check() error {
	if c.Server == nil {
		return errors.New("invalid server config")
	}

	if c.Redis == nil {
		return errors.New("invalid redis config")
	}

	if c.Queue == nil {
		return errors.New("invalid queue config")
	}

	if len(c.Queue.GroupName) == 0 {
		return errors.New("invalid group name")
	}
	return nil
}
