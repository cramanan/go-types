package booleans

func (b Boolean) ToInt() int {
	return ToInt(b)
}

func (b Boolean) AND(b2 Boolean) Boolean {
	return AND(b, b2)
}

func (b Boolean) OR(b2 Boolean) Boolean {
	return OR(b, b2)
}

func (b Boolean) NOT() Boolean {
	return NOT(b)
}

func (b Boolean) NAND(b2 Boolean) Boolean {
	return NAND(b, b2)
}

func (b Boolean) NOR(b2 Boolean) Boolean {
	return NOR(b, b2)
}
