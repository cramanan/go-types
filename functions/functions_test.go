package functions_test

import (
	"testing"

	. "github.com/cramanan/go-types/functions"
)

func TestFunctions(t *testing.T) {
	testCases := []struct {
		desc      string
		got, want bool
	}{
		{"Equal", Equal(1, 1), true},
		{"Equal", Equal(0, 1), false},

		{"Greater", Greater(1, 0), true},
		{"Greater", Greater(0, 0), false},
		{"Greater", Greater(0, 1), false},

		{"GreaterOrEqual", GreaterOrEqual(1, 0), true},
		{"GreaterOrEqual", GreaterOrEqual(0, 0), true},
		{"GreaterOrEqual", GreaterOrEqual(0, 1), false},

		{"Less", Less(0, 1), true},
		{"Less", Less(0, 0), false},
		{"Less", Less(1, 0), false},

		{"LessOrEqual", LessOrEqual(0, 1), true},
		{"LessOrEqual", LessOrEqual(0, 0), true},
		{"LessOrEqual", LessOrEqual(1, 0), false},

		{"Not Equal", NotEqual(1, 0), true},
		{"Not Equal", NotEqual(1, 1), false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("%s got %t, want %t", tC.desc, tC.got, tC.want)
			}
		})
	}
}
