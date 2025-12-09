package main

import "testing"

func TestCountTriples(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1 - n=5",
			n:    5,
			want: 2,
		},
		{
			name: "example 2 - n=10",
			n:    10,
			want: 4,
		},
		{
			name: "minimum value - n=4",
			n:    4,
			want: 0,
		},
		{
			name: "n=13",
			n:    13,
			want: 6,
		},
		{
			name: "n=15",
			n:    15,
			want: 8,
		},
		{
			name: "n=25",
			n:    25,
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTriples(tt.n)
			if got != tt.want {
				t.Errorf("countTriples(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkCountTriples(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=5", 5},
		{"n=10", 10},
		{"n=25", 25},
		{"n=50", 50},
		{"n=100", 100},
		{"n=250", 250},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countTriples(bm.n)
			}
		})
	}
}
