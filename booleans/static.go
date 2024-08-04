package booleans

// NOT returns the logical negation of b.
func NOT(b Boolean) Boolean { return !b }

// AND returns the logical conjunction of a and b.
func AND(a, b Boolean) Boolean { return a && b }

// NAND returns the logical negation of the conjunction of a and b.
func NAND(a, b Boolean) Boolean { return !(a && b) }

// OR returns the logical disjunction of a and b.
func OR(a, b Boolean) Boolean { return a || b }

// NOR returns the logical negation of the disjunction of a and b.
func NOR(a, b Boolean) Boolean { return !(a || b) }
