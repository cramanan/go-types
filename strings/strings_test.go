package strings

import "testing"

func TestNew(t *testing.T) {
	got := New()
	want := String("")
	if got != want {
		t.Errorf("New() = %s, want %s", got, want)
	}
}
