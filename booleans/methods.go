package booleans

// AND returns the logical conjunction of a and b.
func (b Boolean) AND(b2 Boolean) Boolean { return AND(b, b2) }

// OR returns the logical disjunction of a and b.
func (b Boolean) OR(b2 Boolean) Boolean { return OR(b, b2) }

// NOT returns the logical negation of b.
func (b Boolean) NOT() Boolean { return NOT(b) }

// NAND returns the logical negation of the conjunction of a and b.
func (b Boolean) NAND(b2 Boolean) Boolean { return NAND(b, b2) }

// NOR returns the logical negation of the disjunction of a and b.
func (b Boolean) NOR(b2 Boolean) Boolean { return NOR(b, b2) }

func (b1 Boolean) XOR(b2 Boolean) Boolean { return XOR(b1, b2) }
