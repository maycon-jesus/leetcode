package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Casos válidos
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single pair of parentheses",
			input:    "()",
			expected: true,
		},
		{
			name:     "multiple pairs",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "nested parentheses",
			input:    "([])",
			expected: true,
		},
		{
			name:     "complex nested structure",
			input:    "{[()]}",
			expected: true,
		},
		{
			name:     "multiple nested pairs",
			input:    "((()))",
			expected: true,
		},
		{
			name:     "mixed nested structure",
			input:    "{[()()]}",
			expected: true,
		},
		{
			name:     "long valid sequence",
			input:    "(){}[](){}[]",
			expected: true,
		},

		// Casos inválidos
		{
			name:     "single opening parenthesis",
			input:    "(",
			expected: false,
		},
		{
			name:     "single closing parenthesis",
			input:    ")",
			expected: false,
		},
		{
			name:     "mismatched pair",
			input:    "(]",
			expected: false,
		},
		{
			name:     "wrong order",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "unclosed opening bracket",
			input:    "{[]",
			expected: false,
		},
		{
			name:     "extra closing bracket",
			input:    "{[]}]",
			expected: false,
		},
		{
			name:     "only opening brackets",
			input:    "((((",
			expected: false,
		},
		{
			name:     "only closing brackets",
			input:    "))))",
			expected: false,
		},
		{
			name:     "odd length string",
			input:    "(()",
			expected: false,
		},
		{
			name:     "closing before opening",
			input:    ")()",
			expected: false,
		},
		{
			name:     "complex invalid structure",
			input:    "{[(])}",
			expected: false,
		},
		{
			name:     "mismatched nested",
			input:    "((})",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{
			name:  "short valid",
			input: "()",
		},
		{
			name:  "medium valid",
			input: "()[]{}",
		},
		{
			name:  "nested valid",
			input: "{[()()]}",
		},
		{
			name:  "long valid",
			input: "(){}[](){}[](){}[]",
		},
		{
			name:  "invalid early detection",
			input: ")()",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isValid(bm.input)
			}
		})
	}
}
