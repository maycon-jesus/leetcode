package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		// Examples from problem description
		{
			name:     "Example 1: flower, flow, flight",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Example 2: No common prefix",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		// Single string cases
		{
			name:     "Single string",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "Single empty string",
			input:    []string{""},
			expected: "",
		},
		// Identical strings
		{
			name:     "All strings identical",
			input:    []string{"test", "test", "test"},
			expected: "test",
		},
		{
			name:     "Two identical strings",
			input:    []string{"abc", "abc"},
			expected: "abc",
		},
		// Prefix cases
		{
			name:     "One string is prefix of others",
			input:    []string{"prefix", "prefixes", "prefixing"},
			expected: "prefix",
		},
		{
			name:     "First string is shortest prefix",
			input:    []string{"a", "abc", "abcd"},
			expected: "a",
		},
		{
			name:     "Last string is shortest prefix",
			input:    []string{"abcd", "abc", "a"},
			expected: "a",
		},
		{
			name:     "Middle string is shortest prefix",
			input:    []string{"abcd", "a", "abc"},
			expected: "a",
		},
		// Empty string cases
		{
			name:     "Empty string at beginning",
			input:    []string{"", "b"},
			expected: "",
		},
		{
			name:     "Empty string in middle",
			input:    []string{"abc", "", "abd"},
			expected: "",
		},
		{
			name:     "Empty string at end",
			input:    []string{"abc", "abd", ""},
			expected: "",
		},
		{
			name:     "All empty strings",
			input:    []string{"", "", ""},
			expected: "",
		},
		// No common prefix
		{
			name:     "Different first characters",
			input:    []string{"apple", "banana", "cherry"},
			expected: "",
		},
		{
			name:     "Common first char only",
			input:    []string{"a", "ab", "ac"},
			expected: "a",
		},
		// Long common prefix
		{
			name:     "Long common prefix",
			input:    []string{"interspecies", "interstellar", "interstate"},
			expected: "inters",
		},
		{
			name:     "Very long identical prefix",
			input:    []string{"abcdefghijklmnop", "abcdefghijklmnopq", "abcdefghijklmnopqr"},
			expected: "abcdefghijklmnop",
		},
		// Two strings
		{
			name:     "Two strings with common prefix",
			input:    []string{"hello", "help"},
			expected: "hel",
		},
		{
			name:     "Two strings no common prefix",
			input:    []string{"abc", "xyz"},
			expected: "",
		},
		// Edge cases
		{
			name:     "Single character strings - same",
			input:    []string{"a", "a", "a"},
			expected: "a",
		},
		{
			name:     "Single character strings - different",
			input:    []string{"a", "b", "c"},
			expected: "",
		},
		{
			name:     "Prefix at different positions",
			input:    []string{"abc", "ab", "abcd"},
			expected: "ab",
		},
		// Real-world examples
		{
			name:     "Programming languages",
			input:    []string{"javascript", "java", "javadoc"},
			expected: "java",
		},
		{
			name:     "File extensions",
			input:    []string{"file.txt", "file.doc", "file.pdf"},
			expected: "file.",
		},
		{
			name:     "URLs with common domain",
			input:    []string{"https://example.com/path1", "https://example.com/path2", "https://example.com/path3"},
			expected: "https://example.com/path",
		},
		// Incremental differences
		{
			name:     "Incremental string lengths",
			input:    []string{"a", "ab", "abc", "abcd", "abcde"},
			expected: "a",
		},
		{
			name:     "All share complete first string",
			input:    []string{"test", "testing", "tester", "tested"},
			expected: "test",
		},
		// Special patterns
		{
			name:     "Repeated characters",
			input:    []string{"aaaa", "aaab", "aaac"},
			expected: "aaa",
		},
		{
			name:     "Numbers as strings",
			input:    []string{"123456", "123789", "123000"},
			expected: "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkLongestCommonPrefix(b *testing.B) {
	benchmarks := []struct {
		name  string
		input []string
	}{
		{
			name:  "Short strings with prefix",
			input: []string{"flower", "flow", "flight"},
		},
		{
			name:  "No common prefix",
			input: []string{"dog", "racecar", "car"},
		},
		{
			name:  "Long common prefix",
			input: []string{"interspecies", "interstellar", "interstate"},
		},
		{
			name:  "Many identical strings",
			input: []string{"test", "test", "test", "test", "test", "test", "test", "test"},
		},
		{
			name:  "Early mismatch",
			input: []string{"abc", "xyz", "def"},
		},
		{
			name:  "Large array with short prefix",
			input: []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "a10"},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				longestCommonPrefix(bm.input)
			}
		})
	}
}

// Benchmark memory allocations
func BenchmarkLongestCommonPrefixMemory(b *testing.B) {
	input := []string{"flower", "flow", "flight"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		longestCommonPrefix(input)
	}
}
