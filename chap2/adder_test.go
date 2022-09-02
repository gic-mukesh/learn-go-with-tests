package integers

import "testing"

func TestAdder(t *testing.T) {
	got := sum(2, 2)
	want := 4

	if got != want {
		t.Errorf("expected %d but got %d", want, got)
	}
}
