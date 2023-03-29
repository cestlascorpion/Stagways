package core

import (
	"time"
)

const (
	StreamMaxLen = 4096
	StreamExpire = time.Hour * 24
	StreamLimit  = 128
	StreamBlock  = time.Second * 2
)
