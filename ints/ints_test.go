package ints

import (
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	want := Int(0)
	if got != want {
		t.Errorf("New() = %d, want %d", got, want)
	}
}

func TestInt_String(t *testing.T) {
	tests := []struct {
		i    Int
		want string
	}{
		{Int(math.MaxInt64), "9223372036854775807"},
		{Int(math.MinInt64), "-9223372036854775808"},
		{Int(0), "0"},
		{Int(123), "123"},
		{Int(-123), "-123"},
		{Int(1234567890), "1234567890"},
		{Int(-1234567890), "-1234567890"},
	}

	for _, tt := range tests {
		if got := tt.i.String(); got != tt.want {
			t.Errorf("Int(%d).String() = %q, want %q", tt.i, got, tt.want)
		}
	}
}
