package booleans_test

import (
	"testing"

	. "github.com/cramanan/go-types/booleans"
)

func TestStaticLogical(t *testing.T) {
	testCases := []struct {
		desc      string
		got, want Boolean
	}{
		{"AND", AND(True, True), True},
		{"AND", AND(False, True), False},
		{"AND", AND(False, False), False},
		{"AND", AND(True, False), False},

		{"OR", OR(True, True), True},
		{"OR", OR(False, True), True},
		{"OR", OR(False, False), False},
		{"OR", OR(True, False), True},

		{"NOT", NOT(True), False},
		{"NOT", NOT(False), True},

		{"NAND", NAND(True, True), False},
		{"NAND", NAND(False, True), True},
		{"NAND", NAND(False, False), True},
		{"NAND", NAND(True, False), True},

		{"NOR", NOR(True, True), False},
		{"NOR", NOR(False, True), False},
		{"NOR", NOR(False, False), True},
		{"NOR", NOR(True, False), False},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("%s got %t, want %t", tC.desc, tC.got, tC.want)
			}
		})
	}
}
