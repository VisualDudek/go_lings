// Package main demonstrates computed constants using iota and expressions.
// Learn: Constants can be derived from other constants; iota increments in const blocks; untyped constants maintain precision;
// bit operations create flags and enumerations; constants adapt to context types.
package main

// Primary USe Case: enumeration with `iota`

const (
	StatusOK        = iota + 200 // 200
	StatusCreated                // 201
	StatusAccepted               // 202
	StatusNoContent              // 203
)

// Mathematical Conversions Constants
const (
	Byte = 1
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
	TB   = 1024 * GB
)

const (
	PowerOf2_8  = 1 << 8
	PowerOf2_16 = 1 << 16
	MaxUint16   = PowerOf2_16 - 1
)

// The Untyped Constant Advantage
const (
	HighPrecision = 1e20
	LowPrecision  = HighPrecision / 1e10

	computed = 42
)

// Bit Flags & Permissions System
const (
	READ    = 1 << iota // 1 (0b001)
	WRITE               // 2 (0b010)
	EXECUTE             // 4 (0b100)
)

// Combination of Permissions
const (
	READ_WRITE = READ | WRITE           // 3 (0b011)
	ALL        = READ | WRITE | EXECUTE // 7 (0b111)
)

func main() {
	userPerms := READ | EXECUTE // User has read and execute permissions

	if userPerms&READ != 0 {
		// User can read
	}

	if userPerms&WRITE == 0 {
		// User cannot write
	}

	x := computed          // can be int
	y := float64(computed) // can be float64
	z := uint8(computed)   // can be uint8
	println(x, y, z)

}
