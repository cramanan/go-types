package ints

type Int int

func New() Int {
	return Int(0)
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func From[I Ordered](i I) Int {
	return Int(i)
}

func (i Int) String() (s string) {
	sign := ""

	if i == -9223372036854775808 {
		return "-9223372036854775808"
	}

	if i == 0 {
		return "0"
	}

	if i < 0 {
		sign = "-"
		i = -i
	}

	for i > 0 {
		s = string(byte('0'+i%10)) + s
		i /= 10
	}

	return sign + s
}

func (i Int) IsNegative() bool {
	return i < 0
}

func (i Int) IsPositive() bool {
	return i > 0
}
