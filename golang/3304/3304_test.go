package main

import "testing"

func TestKthCharacter(t *testing.T) {
	tests := []struct {
		k    int
		want byte
	}{
		{5, 'b'},  // Example 1
		{10, 'c'}, // Example 2
	}

	for _, tt := range tests {
		t.Run("kthCharacter", func(t *testing.T) {
			got := kthCharacter(tt.k)
			if got != tt.want {
				t.Errorf("kthCharacter(%d) = %c; want %c", tt.k, got, tt.want)
			}
		})
	}
}
