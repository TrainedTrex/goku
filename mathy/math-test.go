package mathy

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{
				name:     "positive numbers",
				input:    []int{5, 3, 7, 1, 9},
				expected: 1,
			},
			{
				name:     "negative numbers",
				input:    []int{-5, -3, -7, -1, -9},
				expected: -9,
			},
			{
				name:     "mixed numbers",
				input:    []int{-5, 3, -7, 1, 9},
				expected: -7,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Min(tt.input...); got != tt.expected {
					t.Errorf("Min() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test different integer types
		if got := Min(int8(4), int8(2), int8(6)); got != int8(2) {
			t.Errorf("Min(int8) = %v, want %v", got, int8(2))
		}

		if got := Min(int16(400), int16(200), int16(600)); got != int16(200) {
			t.Errorf("Min(int16) = %v, want %v", got, int16(200))
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []float64
			expected float64
		}{
			{
				name:     "positive numbers",
				input:    []float64{5.5, 3.3, 7.7, 1.1, 9.9},
				expected: 1.1,
			},
			{
				name:     "negative numbers",
				input:    []float64{-5.5, -3.3, -7.7, -1.1, -9.9},
				expected: -9.9,
			},
			{
				name:     "mixed numbers",
				input:    []float64{-5.5, 3.3, -7.7, 1.1, 9.9},
				expected: -7.7,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Min(tt.input...); got != tt.expected {
					t.Errorf("Min() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test float32
		if got := Min(float32(3.14), float32(2.71), float32(1.41)); got != float32(1.41) {
			t.Errorf("Min(float32) = %v, want %v", got, float32(1.41))
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{
				name:     "positive numbers",
				input:    []int{5, 3, 7, 1, 9},
				expected: 9,
			},
			{
				name:     "negative numbers",
				input:    []int{-5, -3, -7, -1, -9},
				expected: -1,
			},
			{
				name:     "mixed numbers",
				input:    []int{-5, 3, -7, 1, 9},
				expected: 9,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Max(tt.input...); got != tt.expected {
					t.Errorf("Max() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test different integer types
		if got := Max(int8(4), int8(2), int8(6)); got != int8(6) {
			t.Errorf("Max(int8) = %v, want %v", got, int8(6))
		}

		if got := Max(int16(400), int16(200), int16(600)); got != int16(600) {
			t.Errorf("Max(int16) = %v, want %v", got, int16(600))
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []float64
			expected float64
		}{
			{
				name:     "positive numbers",
				input:    []float64{5.5, 3.3, 7.7, 1.1, 9.9},
				expected: 9.9,
			},
			{
				name:     "negative numbers",
				input:    []float64{-5.5, -3.3, -7.7, -1.1, -9.9},
				expected: -1.1,
			},
			{
				name:     "mixed numbers",
				input:    []float64{-5.5, 3.3, -7.7, 1.1, 9.9},
				expected: 9.9,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Max(tt.input...); got != tt.expected {
					t.Errorf("Max() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test float32
		if got := Max(float32(3.14), float32(2.71), float32(1.41)); got != float32(3.14) {
			t.Errorf("Max(float32) = %v, want %v", got, float32(3.14))
		}
	})
}

func TestAbs(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    int
			expected int
		}{
			{
				name:     "positive number",
				input:    5,
				expected: 5,
			},
			{
				name:     "negative number",
				input:    -5,
				expected: 5,
			},
			{
				name:     "zero",
				input:    0,
				expected: 0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Abs(tt.input); got != tt.expected {
					t.Errorf("Abs() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test different integer types
		if got := Abs(int8(-8)); got != int8(8) {
			t.Errorf("Abs(int8) = %v, want %v", got, int8(8))
		}

		if got := Abs(int16(-16)); got != int16(16) {
			t.Errorf("Abs(int16) = %v, want %v", got, int16(16))
		}

		if got := Abs(int32(-32)); got != int32(32) {
			t.Errorf("Abs(int32) = %v, want %v", got, int32(32))
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			name     string
			input    float64
			expected float64
		}{
			{
				name:     "positive float",
				input:    3.14,
				expected: 3.14,
			},
			{
				name:     "negative float",
				input:    -3.14,
				expected: 3.14,
			},
			{
				name:     "zero float",
				input:    0.0,
				expected: 0.0,
			},
			{
				name:     "small negative float",
				input:    -0.00001,
				expected: 0.00001,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Abs(tt.input); got != tt.expected {
					t.Errorf("Abs() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test float32
		if got := Abs(float32(-3.14)); got != float32(3.14) {
			t.Errorf("Abs(float32) = %v, want %v", got, float32(3.14))
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{
				name:     "positive numbers",
				input:    []int{1, 2, 3, 4, 5},
				expected: 15,
			},
			{
				name:     "negative numbers",
				input:    []int{-1, -2, -3, -4, -5},
				expected: -15,
			},
			{
				name:     "mixed numbers",
				input:    []int{-1, 2, -3, 4, -5},
				expected: -3,
			},
			{
				name:     "empty slice",
				input:    []int{},
				expected: 0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Sum(tt.input); got != tt.expected {
					t.Errorf("Sum() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test different integer types
		if got := Sum([]int8{1, 2, 3}); got != int8(6) {
			t.Errorf("Sum(int8) = %v, want %v", got, int8(6))
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []float64
			expected float64
		}{
			{
				name:     "positive numbers",
				input:    []float64{1.1, 2.2, 3.3},
				expected: 6.6,
			},
			{
				name:     "negative numbers",
				input:    []float64{-1.1, -2.2, -3.3},
				expected: -6.6,
			},
			{
				name:     "mixed numbers",
				input:    []float64{-1.1, 2.2, -3.3},
				expected: -2.2,
			},
			{
				name:     "empty slice",
				input:    []float64{},
				expected: 0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Sum(tt.input); math.Abs(got-tt.expected) > 0.0001 {
					t.Errorf("Sum() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test float32
		float32Slice := []float32{1.1, 2.2, 3.3}
		if got := Sum(float32Slice); math.Abs(float64(got)-6.6) > 0.0001 {
			t.Errorf("Sum(float32) = %v, want %v", got, 6.6)
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected int
		}{
			{
				name:     "positive numbers",
				input:    []int{1, 2, 3, 4, 5},
				expected: 120,
			},
			{
				name:     "with zero",
				input:    []int{1, 2, 0, 4, 5},
				expected: 0,
			},
			{
				name:     "negative numbers",
				input:    []int{-1, -2, -3},
				expected: -6,
			},
			{
				name:     "empty slice",
				input:    []int{},
				expected: 1,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Multiply(tt.input); got != tt.expected {
					t.Errorf("Multiply() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test different integer types
		if got := Multiply([]int8{1, 2, 3}); got != int8(6) {
			t.Errorf("Multiply(int8) = %v, want %v", got, int8(6))
		}
	})

	t.Run("floats", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []float64
			expected float64
		}{
			{
				name:     "positive numbers",
				input:    []float64{1.1, 2.2, 3.3},
				expected: 7.986,
			},
			{
				name:     "with zero",
				input:    []float64{1.1, 0.0, 3.3},
				expected: 0.0,
			},
			{
				name:     "negative numbers",
				input:    []float64{-1.1, -2.2},
				expected: 2.42,
			},
			{
				name:     "empty slice",
				input:    []float64{},
				expected: 1.0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Multiply(tt.input); math.Abs(got-tt.expected) > 0.0001 {
					t.Errorf("Multiply() = %v, want %v", got, tt.expected)
				}
			})
		}

		// Test float32
		float32Slice := []float32{1.1, 2.2, 3.3}
		if got := Multiply(float32Slice); math.Abs(float64(got)-7.986) > 0.0001 {
			t.Errorf("Multiply(float32) = %v, want %v", got, 7.986)
		}
	})
}
