package main

import (
	"errors"
	"testing"
	"time"
)

func TestTry_noError(t *testing.T) {
	r := &Retry{}

	err := Try(func() error {
		return nil
	}, r)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint(1)
	if r.Attempt != expected {
		t.Errorf("expected count to be %d got %d", expected, r.Attempt)
	}
}

func TestTry_error(t *testing.T) {
	r := &Retry{Max: 1}

	err := Try(func() error {
		return errors.New("error")
	}, r)
	if err == nil {
		t.Error("expected error")
	}

	expected := uint(2)
	if r.Attempt != expected {
		t.Errorf("expected count to be %d got %d", expected, r.Attempt)
	}
}

func TestTry_timeout(t *testing.T) {
	r := &Retry{Max: 1, Timeout: 2 * time.Millisecond}

	start := time.Now()
	err := Try(func() error {
		return errors.New("error")
	}, r)
	if err == nil {
		t.Error("expected error")
	}

	diff := time.Now().Sub(start)
	if diff < r.Timeout {
		t.Errorf("expected %s to be less than %s", r.Timeout, diff)
	}
}
