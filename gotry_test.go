package main

import (
	"errors"
	"testing"
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
	r := &Retry{Max: 2}

	err := Try(func() error {
		return errors.New("error")
	}, r)
	if err == nil {
		t.Error("expected error")
	}

	if r.Attempt != r.Max {
		t.Errorf("expected count to be %d got %d", r.Max, r.Attempt)
	}
}
