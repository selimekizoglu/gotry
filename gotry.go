package main

func Try(f func() error, r *Retry) error {
	err := f()
	r.Attempt++
	for err != nil && r.Attempt < r.Max {
		err = f()
		r.Attempt++
	}
	return err
}
