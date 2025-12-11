package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		// Basic single symbols
		{
			name:     "Single I",
			input:    "I",
			expected: 1,
		},
		{
			name:     "Single V",
			input:    "V",
			expected: 5,
		},
		{
			name:     "Single X",
			input:    "X",
			expected: 10,
		},
		{
			name:     "Single L",
			input:    "L",
			expected: 50,
		},
		{
			name:     "Single C",
			input:    "C",
			expected: 100,
		},
		{
			name:     "Single D",
			input:    "D",
			expected: 500,
		},
		{
			name:     "Single M",
			input:    "M",
			expected: 1000,
		},
		// Examples from problem description
		{
			name:     "Example 1: III",
			input:    "III",
			expected: 3,
		},
		{
			name:     "Example 2: LVIII",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "Example 3: MCMXCIV",
			input:    "MCMXCIV",
			expected: 1994,
		},
		// Subtractive cases
		{
			name:     "IV (4)",
			input:    "IV",
			expected: 4,
		},
		{
			name:     "IX (9)",
			input:    "IX",
			expected: 9,
		},
		{
			name:     "XL (40)",
			input:    "XL",
			expected: 40,
		},
		{
			name:     "XC (90)",
			input:    "XC",
			expected: 90,
		},
		{
			name:     "CD (400)",
			input:    "CD",
			expected: 400,
		},
		{
			name:     "CM (900)",
			input:    "CM",
			expected: 900,
		},
		// Complex combinations
		{
			name:     "XLII (42)",
			input:    "XLII",
			expected: 42,
		},
		{
			name:     "XCIV (94)",
			input:    "XCIV",
			expected: 94,
		},
		{
			name:     "CDXLIV (444)",
			input:    "CDXLIV",
			expected: 444,
		},
		{
			name:     "CMXCIX (999)",
			input:    "CMXCIX",
			expected: 999,
		},
		{
			name:     "MMMCMXCIX (3999)",
			input:    "MMMCMXCIX",
			expected: 3999,
		},
		// Additional test cases
		{
			name:     "II (2)",
			input:    "II",
			expected: 2,
		},
		{
			name:     "XII (12)",
			input:    "XII",
			expected: 12,
		},
		{
			name:     "XXVII (27)",
			input:    "XXVII",
			expected: 27,
		},
		{
			name:     "MMXX (2020)",
			input:    "MMXX",
			expected: 2020,
		},
		{
			name:     "MMXXIV (2024)",
			input:    "MMXXIV",
			expected: 2024,
		},
		{
			name:     "DCCCXC (890)",
			input:    "DCCCXC",
			expected: 890,
		},
		{
			name:     "MCDXLIV (1444)",
			input:    "MCDXLIV",
			expected: 1444,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInt(tt.input)
			if result != tt.expected {
				t.Errorf("romanToInt(%s) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkRomanToInt(b *testing.B) {
	testCases := []string{
		"III",
		"LVIII",
		"MCMXCIV",
		"MMMCMXCIX",
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				romanToInt(tc)
			}
		})
	}
}
