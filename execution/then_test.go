package execution

import (
	"testing"
	"time"
)

func TestThen(t *testing.T) {
	wanted := 100
	s := Then(Just(10), func(v int) int {
		return v * v
	})
	if v, err := SyncWait(s); err != nil {
		t.Errorf("expected no error, got error: '%v'", err)
	} else if v != wanted {
		t.Errorf("expected v=%d, got v=%d", wanted, v)
	}
}

func TestThenStopped(t *testing.T) {
	s := Then(Just(10), func(v int) int {
		time.Sleep(2 * time.Second)
		return v * v
	})

	go func() {
		time.Sleep(1 * time.Second)
		// Send the stop signal
		s.Stop()
	}()

	start := time.Now()
	if _, err := SyncWait(s); err != nil {
		t.Errorf("expected no error, got error: '%v'", err)
	}
	duration := time.Now().Sub(start)

	if !s.Stopped() {
		t.Errorf("expected execution chain to be stopped")
	}

	if duration >= time.Second*2 {
		t.Errorf("expected execution chain to stop after one second, but stopped after: %.02f seconds",
			duration.Seconds())
	}
}
