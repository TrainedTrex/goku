package mathy

import "golang.org/x/exp/constraints"

// LCM calculates the Least Common Multiple of two integers.
// Note: Only works with integer types as LCM is not defined for floating-point numbers.
func LCM[T constraints.Signed](a, b T) T {
	// Handle special cases to prevent overflow
	if a == 0 || b == 0 {
		return 0
	}

	// Use absolute values to handle negative numbers
	absA := Abs(a)
	absB := Abs(b)

	// Calculate GCD first
	gcd := GCD(absA, absB)

	// Calculate LCM using the formula: |a * b| / GCD(a,b)
	// Use integer division to avoid floating point
	return (absA / gcd) * absB
}

// GCD calculates the Greatest Common Divisor of two integers using Euclidean algorithm.
// Note: Only works with integer types as GCD is not defined for floating-point numbers.
func GCD[T constraints.Signed](a, b T) T {
	// Handle special cases
	if a == 0 {
		return Abs(b)
	}
	if b == 0 {
		return Abs(a)
	}

	// Use absolute values to handle negative numbers
	a = Abs(a)
	b = Abs(b)

	// Ensure a is the larger number
	if b > a {
		a, b = b, a
	}

	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
