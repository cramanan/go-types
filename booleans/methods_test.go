package booleans

import "testing"

func TestLogical(t *testing.T) {
	testCases := []struct {
		desc      string
		got, want Boolean
	}{
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
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.got != tC.want {
				t.Errorf("%s got %t, want %t", tC.desc, tC.got, tC.want)
			}
		})
	}
}
