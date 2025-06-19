package main

import "testing"

func TestGenerateTag(t *testing.T) {
	tests := []struct {
		name     string
		caption  string
		expected string
	}{
		{
			name:     "simple case",
			caption:  "hello world",
			expected: "#helloWorld",
		},
		{
			name:     "leading and trailing spaces",
			caption:  "  hello world  ",
			expected: "#helloWorld",
		},
		{
			name:     "multiple spaces between words",
			caption:  "hello   world",
			expected: "#helloWorld",
		},
		{
			name:     "mixed case",
			caption:  "Hello World",
			expected: "#helloWorld",
		},
		{
			name:     "long caption",
			caption:  "a very long caption that will be truncated to ninety nine characters a very long caption that will be truncated to ninety nine characters",
			expected: "#aVeryLongCaptionThatWillBeTruncatedToNinetyNineCharactersAVeryLongCaptionThatWillBeTruncatedToNinety",
		},
		{
			name:     "empty caption",
			caption:  "",
			expected: "#",
		},
		{
			name:     "single word",
			caption:  "word",
			expected: "#word",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := generateTag(tt.caption)
			if actual != tt.expected {
				t.Errorf("expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func BenchmarkGenerateTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateTag("a very long caption that will be truncated to ninety nine characters a very long caption that will be truncated to ninety nine characters")
	}
}
