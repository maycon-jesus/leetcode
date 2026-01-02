package main

import "testing"

func TestMaximumProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		k      int
		want   int64
	}{
		{
			name:   "Example 1: Mix of normal and short selling",
			prices: []int{1, 7, 9, 8, 2},
			k:      2,
			want:   14,
		},
		{
			name:   "Example 2: Three transactions",
			prices: []int{12, 16, 19, 19, 8, 1, 19, 13, 9},
			k:      3,
			want:   36,
		},
		{
			name:   "Only normal transactions profitable",
			prices: []int{1, 5, 2, 9, 3, 8},
			k:      2,
			want:   13, // (5-1) + (9-2) = 4 + 7 = 11 OU (9-1) + (8-3) = 8 + 5 = 13
		},
		{
			name:   "Only short selling profitable",
			prices: []int{10, 5, 8, 3, 7, 2},
			k:      2,
			want:   12, // Short: (10-5) = 5, Long: (8-3) = 5, Total = 10 OU Short: (10-2) = 8 e Long: (8-3) = 5 - melhor: 12
		},
		{
			name:   "Single transaction allowed",
			prices: []int{3, 1, 4, 8, 2},
			k:      1,
			want:   7, // 8-1 = 7
		},
		{
			name:   "Prices always increasing",
			prices: []int{1, 2, 3, 4, 5},
			k:      2,
			want:   4, // (5-1) = 4
		},
		{
			name:   "Prices always decreasing",
			prices: []int{5, 4, 3, 2, 1},
			k:      2,
			want:   4, // Short: (5-1) = 4
		},
		{
			name:   "All same prices",
			prices: []int{5, 5, 5, 5, 5},
			k:      2,
			want:   0, // No profit possible
		},
		{
			name:   "Two elements only",
			prices: []int{1, 5},
			k:      1,
			want:   4,
		},
		{
			name:   "k equals 0",
			prices: []int{1, 5, 3, 8},
			k:      0,
			want:   0,
		},
		{
			name:   "Large price difference",
			prices: []int{1, 1000000000},
			k:      1,
			want:   999999999,
		},
		{
			name:   "Alternating high and low",
			prices: []int{10, 1, 10, 1, 10},
			k:      2,
			want:   18, // (10-1) + (10-1) = 18
		},
		{
			name:   "Complex mixed strategy",
			prices: []int{5, 10, 5, 10, 5},
			k:      2,
			want:   10, // (10-5) + (10-5) = 10
		},
		{
			name:   "Maximum k allowed",
			prices: []int{1, 5, 2, 8, 3, 9},
			k:      3,
			want:   16, // (5-1) + (8-2) + (9-3) = 4 + 6 + 6 = 16
		},
		{
			name:   "Short then long",
			prices: []int{10, 5, 1, 8},
			k:      2,
			want:   12, // Short: (10-5) = 5, Long: (8-1) = 7, Total = 12
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumProfit(tt.prices, tt.k)
			if got != tt.want {
				t.Errorf("maximumProfit(%v, %d) = %d; want %d", tt.prices, tt.k, got, tt.want)
			}
		})
	}
}

func TestMaximumProfitEdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		k      int
		want   int64
	}{
		{
			name:   "Empty prices array",
			prices: []int{},
			k:      1,
			want:   0,
		},
		{
			name:   "Single element",
			prices: []int{5},
			k:      1,
			want:   0,
		},
		{
			name:   "Minimum constraint values",
			prices: []int{1, 1},
			k:      1,
			want:   0,
		},
		{
			name:   "Maximum price value",
			prices: []int{1, 1000000000, 1},
			k:      1,
			want:   999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumProfit(tt.prices, tt.k)
			if got != tt.want {
				t.Errorf("maximumProfit(%v, %d) = %d; want %d", tt.prices, tt.k, got, tt.want)
			}
		})
	}
}

func TestMaxFunction(t *testing.T) {
	tests := []struct {
		name string
		a    int64
		b    int64
		want int64
	}{
		{"a greater", 10, 5, 10},
		{"b greater", 5, 10, 10},
		{"equal", 7, 7, 7},
		{"negative values", -5, -10, -5},
		{"zero and positive", 0, 5, 5},
		{"zero and negative", 0, -5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := max(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkMaximumProfit(b *testing.B) {
	prices := []int{12, 16, 19, 19, 8, 1, 19, 13, 9}
	k := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximumProfit(prices, k)
	}
}

func BenchmarkMaximumProfitLarge(b *testing.B) {
	// Create a larger test case
	prices := make([]int, 1000)
	for i := range prices {
		prices[i] = (i * 7) % 100
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximumProfit(prices, k)
	}
}
