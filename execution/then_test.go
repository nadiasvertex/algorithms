package execution

import "testing"

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
	stopped := false
	s := Then(Just(10), func(v int) int {
		return v * v
	}).OnStopped(func() {
		stopped = true
	})

	if _, err := SyncWait(s); err != nil {
		t.Errorf("expected no error, got error: '%v'", err)
	} else if !stopped {
		t.Errorf("expected execution chain to be stopped")
	}
}
