package gotry

import (
	"time"
)

func Try(f func() error, r Retry) error {
	err := f()
	attempt := 1
	for err != nil && attempt <= r.Max {
		if r.Timeout > 0 {
			timer := time.NewTimer(r.Timeout)
			<-timer.C
		}
		err = f()
		attempt++
	}
	return err
}
