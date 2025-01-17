package mathy

import "testing"

func TestLCM(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        12,
			b:        18,
			expected: 36,
		},
		{
			name:     "one negative number",
			a:        -12,
			b:        18,
			expected: 36,
		},
		{
			name:     "both negative numbers",
			a:        -12,
			b:        -18,
			expected: 36,
		},
		{
			name:     "one zero",
			a:        0,
			b:        5,
			expected: 0,
		},
		{
			name:     "both zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "prime numbers",
			a:        7,
			b:        13,
			expected: 91,
		},
		{
			name:     "same numbers",
			a:        7,
			b:        7,
			expected: 7,
		},
		{
			name:     "one is multiple of other",
			a:        6,
			b:        24,
			expected: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.a, tt.b); got != tt.expected {
				t.Errorf("LCM(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
			// Test commutativity: LCM(a,b) = LCM(b,a)
			if got := LCM(tt.b, tt.a); got != tt.expected {
				t.Errorf("LCM(%v, %v) = %v, want %v", tt.b, tt.a, got, tt.expected)
			}
		})
	}

	// Test different integer types
	t.Run("different integer types", func(t *testing.T) {
		if got := LCM(int8(4), int8(6)); got != int8(12) {
			t.Errorf("LCM(int8) = %v, want %v", got, int8(12))
		}

		if got := LCM(int16(12), int16(18)); got != int16(36) {
			t.Errorf("LCM(int16) = %v, want %v", got, int16(36))
		}

		if got := LCM(int32(15), int32(25)); got != int32(75) {
			t.Errorf("LCM(int32) = %v, want %v", got, int32(75))
		}
	})
}

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        48,
			b:        18,
			expected: 6,
		},
		{
			name:     "one negative number",
			a:        -48,
			b:        18,
			expected: 6,
		},
		{
			name:     "both negative numbers",
			a:        -48,
			b:        -18,
			expected: 6,
		},
		{
			name:     "one zero",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "both zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "prime numbers",
			a:        7,
			b:        13,
			expected: 1,
		},
		{
			name:     "same numbers",
			a:        7,
			b:        7,
			expected: 7,
		},
		{
			name:     "one is multiple of other",
			a:        6,
			b:        24,
			expected: 6,
		},
		{
			name:     "large numbers",
			a:        1071,
			b:        462,
			expected: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.a, tt.b); got != tt.expected {
				t.Errorf("GCD(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
			// Test commutativity: GCD(a,b) = GCD(b,a)
			if got := GCD(tt.b, tt.a); got != tt.expected {
				t.Errorf("GCD(%v, %v) = %v, want %v", tt.b, tt.a, got, tt.expected)
			}
		})
	}

	// Test different integer types
	t.Run("different integer types", func(t *testing.T) {
		if got := GCD(int8(8), int8(12)); got != int8(4) {
			t.Errorf("GCD(int8) = %v, want %v", got, int8(4))
		}

		if got := GCD(int16(48), int16(18)); got != int16(6) {
			t.Errorf("GCD(int16) = %v, want %v", got, int16(6))
		}

		if got := GCD(int32(100), int32(75)); got != int32(25) {
			t.Errorf("GCD(int32) = %v, want %v", got, int32(25))
		}
	})
}
