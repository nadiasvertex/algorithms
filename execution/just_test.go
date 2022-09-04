package execution

import "testing"

func TestJust(t *testing.T) {
	if v, err := SyncWait(Just(10)); err != nil {
		t.Errorf("expected no error, got error: '%v'", err)
	} else if v != 10 {
		t.Errorf("expected v=10, got v=%d", v)
	}
}
