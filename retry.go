package gotry

import (
	"time"
)

type Retry struct {
	Max     int
	Timeout time.Duration
}
