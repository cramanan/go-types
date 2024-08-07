package booleans_test

import (
	"testing"

	. "github.com/cramanan/go-types/booleans"
)

func TestBooleans(t *testing.T) {

	ptr := 1

	testCases := []struct {
		desc      string
		got, want Boolean
	}{
		{"Calibrating", True, true},
		{"Calibrating", False, false},

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

		{"XOR", XOR(False, False), False},
		{"XOR", XOR(True, False), True},
		{"XOR", XOR(False, True), True},
		{"XOR", XOR(True, True), False},

		{"AND", True.AND(True), True},
		{"AND", False.AND(True), False},
		{"AND", False.AND(False), False},
		{"AND", True.AND(False), False},

		{"OR", True.OR(True), True},
		{"OR", False.OR(True), True},
		{"OR", False.OR(False), False},
		{"OR", True.OR(False), True},

		{"NOT", True.NOT(), False},
		{"NOT", False.NOT(), True},

		{"NAND", True.NAND(True), False},
		{"NAND", False.NAND(True), True},
		{"NAND", False.NAND(False), True},
		{"NAND", True.NAND(False), True},

		{"NOR", True.NOR(True), False},
		{"NOR", False.NOR(True), False},
		{"NOR", False.NOR(False), True},
		{"NOR", True.NOR(False), False},

		{"XOR", False.XOR(False), False},
		{"XOR", True.XOR(False), True},
		{"XOR", False.XOR(True), True},
		{"XOR", True.XOR(True), False},

		{"Truthy", Boolean(IsTruthy(True)), True},
		{"Falsy", Boolean(IsTruthy(False)), False},
		{"Falsy", Boolean(IsTruthy(0)), False},
		{"Falsy", Boolean(IsTruthy("")), false},
		{"Truthy", Boolean(IsTruthy("\x00")), True},
		{"Falsy", Boolean(IsTruthy[*int](nil)), false},
		{"Truthy", Boolean(IsTruthy(&ptr)), true},

		{"Truthy", Boolean(IsTruthy(struct{}{})), false},
		{"Truthy", Boolean(IsTruthy(struct{ test string }{"1"})), true},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("%s got %t, want %t", tC.desc, tC.got, tC.want)
			}
		})
	}
}
