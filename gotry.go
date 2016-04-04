package gotry

import (
	"time"
)

func Try(f func() error, r *Retry) error {
	err := f()
	r.Attempt++
	for err != nil && r.Attempt <= r.Max {
		if r.Timeout > 0 {
			timer := time.NewTimer(r.Timeout)
			<-timer.C
		}
		err = f()
		r.Attempt++
	}
	return err
}
