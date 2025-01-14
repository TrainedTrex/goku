package mathy

import "golang.org/x/exp/constraints"

// Number represents any numeric type
type Number interface {
	constraints.Integer | constraints.Float
}

// Min returns the smallest value from the provided numbers
func Min[T Number](nums ...T) T {
	minNum := nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}

// Max returns the largest value from the provided numbers
func Max[T Number](nums ...T) T {
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

// Abs returns the absolute value of the input number
func Abs[T constraints.Signed](in T) T {
	if in < 0 {
		return -in
	}
	return in
}

// Sum returns the summation of all numbers in the slice
func Sum[T Number](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Multiply returns the product of all numbers in the slice
func Multiply[T Number](nums []T) T {
	product := T(1)
	for _, n := range nums {
		product *= n
	}
	return product
}
