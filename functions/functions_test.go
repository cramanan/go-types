package functions_test

import (
	"strings"
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

var longstring = "longStrinGwitHmixofsmaLLandcAps"

func Test_Funcs(t *testing.T) {
	type TestCase[T any] struct {
		desc      string
		got, want T
	}
	testCases := []TestCase[int]{
		{"strings.IndexFunc",
			strings.IndexFunc(longstring, Satisfy('w')),
			strings.IndexRune(longstring, 'w'),
		},
		{"strings.IndexFunc",
			strings.IndexFunc(longstring, Satisfy('s')),
			strings.IndexRune(longstring, 's'),
		},
		{"strings.IndexFunc",
			strings.IndexFunc(longstring, Satisfy('L')),
			strings.IndexRune(longstring, 'L'),
		},
		{"strings.IndexFunc",
			strings.IndexFunc(longstring, Satisfy('X')),
			strings.IndexRune(longstring, 'X'),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("%s got %d, want %d", tC.desc, tC.got, tC.want)
			}
		})
	}
}
