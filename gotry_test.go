package gotry

import (
	"errors"
	"testing"
	"time"
)

func TestTry_noError(t *testing.T) {
	r := Retry{}

	attempt := 0
	err := Try(func() error {
		attempt++
		return nil
	}, r)
	if err != nil {
		t.Fatal(err)
	}

	expected := 1
	if attempt != expected {
		t.Errorf("expected attempt to be %d got %d", expected, attempt)
	}
}

func TestTry_error(t *testing.T) {
	r := Retry{Max: 1}

	attempt := 0
	err := Try(func() error {
		attempt++
		return errors.New("error")
	}, r)
	if err == nil {
		t.Error("expected error")
	}

	expected := 2
	if attempt != expected {
		t.Errorf("expected attempt to be %d got %d", expected, attempt)
	}
}

func TestTry_timeout(t *testing.T) {
	r := Retry{Max: 1, Timeout: 2 * time.Millisecond}

	start := time.Now()
	attempt := 0
	err := Try(func() error {
		attempt++
		return errors.New("error")
	}, r)
	if err == nil {
		t.Error("expected error")
	}

	expected := 2
	if attempt != expected {
		t.Errorf("expected attempt to be %d got %d", expected, attempt)
	}

	diff := time.Now().Sub(start)
	if diff < r.Timeout {
		t.Errorf("expected %s to be less than %s", r.Timeout, diff)
	}
}
