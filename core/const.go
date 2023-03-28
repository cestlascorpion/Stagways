package core

import (
	"errors"
	"time"
)

const (
	StreamMaxLen = 4096
	StreamTTL    = time.Hour * 24
	StreamLimit  = 128
)

var (
	ErrEmptyStream = errors.New("empty stream")
	ErrEmptyData   = errors.New("empty data")
)
