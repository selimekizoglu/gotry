package main

import (
	"time"
)

type Retry struct {
	Max     uint
	Attempt uint
	Timeout time.Duration
}
