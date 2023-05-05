package mq

import (
	"context"
	"sync"
	"time"

	"github.com/cestlascorpion/Stagways/core"
	log "github.com/sirupsen/logrus"
)

type consumerManager struct {
	config  *core.Config
	handler func(ctx context.Context, msg map[string]interface{}, offset string)
	workers map[string]*worker
	mutex   sync.RWMutex
}

func NewConsumerManager(ctx context.Context, config *core.Config, handler func(ctx context.Context, msg map[string]interface{}, offset string)) (Manager, error) {
	return &consumerManager{
		config:  config,
		handler: handler,
		workers: make(map[string]*worker),
	}, nil
}

func (m *consumerManager) AddStream(ctx context.Context, stream string) error {
	if m.exists(stream) {
		log.Debugf("consumer for steam %s exists", stream)
		return nil
	}

	c, err := NewConsumer(ctx, m.config)
	if err != nil {
		log.Errorf("new consumer err %+v", err)
		return err
	}

	err = c.Claim(ctx, stream, m.config.Queue.GroupName)
	if err != nil {
		log.Errorf("consumer claim stream %s err %+v", stream, err)
		return err
	}
	_, err = c.Check(ctx) // not sure about this
	if err != nil {
		log.Errorf("consumer check stream %s err %+v", stream, err)
		return err
	}

	x, cancel := context.WithCancel(context.Background())
	w := &worker{
		consumer: c,
		cancel:   cancel,
	}

	m.startConsumer(x, stream, c)
	m.add(stream, w)
	return nil
}

func (m *consumerManager) DelStream(ctx context.Context, stream string) error {
	m.del(stream)
	return nil
}

func (m *consumerManager) Close(ctx context.Context) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for stream, w := range m.workers {
		w.cancel()
		log.Infof("cancel worker for steam %s", stream)
	}
	return nil
}

func (m *consumerManager) exists(stream string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	_, ok := m.workers[stream]
	return ok
}

func (m *consumerManager) add(stream string, w *worker) {
	m.mutex.Lock()
	_, ok := m.workers[stream]
	if ok {
		m.mutex.Unlock()
		log.Warnf("add exists worker for stream %s", stream)
		w.cancel() // necessary
		return
	}
	m.workers[stream] = w
	m.mutex.Unlock()
	log.Infof("add worker for stream %s", stream)
}

func (m *consumerManager) del(stream string) {
	m.mutex.Lock()
	w, ok := m.workers[stream]
	if !ok {
		m.mutex.Unlock()
		log.Warnf("del no-exists worker for stream %s", stream)
		return
	}
	delete(m.workers, stream)
	m.mutex.Unlock()
	w.cancel() // necessary
	log.Infof("del worker for stream %s", stream)
}

type worker struct {
	consumer Consumer
	cancel   context.CancelFunc
}

func (m *consumerManager) startConsumer(ctx context.Context, stream string, consumer Consumer) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				_ = consumer.Close(context.Background())
				log.Infof("close consumer for stream %s", stream)
				return
			default:
				msg, offset, err := consumer.Consume(ctx)
				if err != nil {
					log.Errorf("consume %s err %+v", stream, err)
					time.Sleep(time.Second)
				} else {
					if m.handler != nil && msg != nil {
						m.handler(ctx, msg, offset)
						log.Debugf("consume %s msg %s ok", stream, offset)
					}
				}
			}
		}
	}()
}
