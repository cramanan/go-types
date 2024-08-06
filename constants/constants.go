package constants

// Byte units

const (
	// B represents a single byte, the basic unit of digital information.
	B float64 = 1 << (10 * iota)

	// KB represents a kilobyte, which is 1,024 bytes.
	KB

	// MB represents a megabyte, which is 1,048,576 bytes or 1,024 kilobytes.
	MB

	// GB represents a gigabyte, which is 1,073,741,824 bytes or 1,024 megabytes.
	GB

	// TB represents a terabyte, which is 1,099,511,627,776 bytes or 1,024 gigabytes.
	TB

	// PB represents a petabyte, which is 1,125,899,906,842,624 bytes or 1,024 terabytes.
	PB

	// EB represents an exabyte, which is 1,152,921,504,606,846,976 bytes or 1,024 petabytes.
	EB

	YB

	ZB
)

// Characters

const (
	NewLine = '\n'
	Tab     = '\t'
	Space   = ' '
)
