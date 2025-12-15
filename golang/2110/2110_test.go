package main

import "testing"

func TestGetDescentPeriods(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int64
	}{
		{
			name:     "Exemplo 1: sequência com quebra",
			prices:   []int{3, 2, 1, 4},
			expected: 7,
		},
		{
			name:     "Exemplo 2: sem períodos descendentes consecutivos",
			prices:   []int{8, 6, 7, 7},
			expected: 4,
		},
		{
			name:     "Exemplo 3: array com um único elemento",
			prices:   []int{1},
			expected: 1,
		},
		{
			name:     "Exemplo 4: sequência completamente descendente",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 15,
		},
		{
			name:     "Exemplo 5: sequência crescente",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Array com dois elementos descendentes",
			prices:   []int{2, 1},
			expected: 3,
		},
		{
			name:     "Array com dois elementos não descendentes",
			prices:   []int{1, 2},
			expected: 2,
		},
		{
			name:     "Múltiplas sequências descendentes",
			prices:   []int{10, 9, 8, 5, 4, 3},
			expected: 12, // [10,9,8] = 6 períodos, [5,4,3] = 6 períodos
		},
		{
			name:     "Sequência com valores iguais",
			prices:   []int{5, 5, 5},
			expected: 3,
		},
		{
			name:     "Sequência longa descendente",
			prices:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 55, // 1+2+3+4+5+6+7+8+9+10 = 55
		},
		{
			name:     "Descidas de tamanho 1 intercaladas",
			prices:   []int{5, 4, 6, 5, 7},
			expected: 7, // 5 elementos individuais + 2 pares descendentes
		},
		{
			name:     "Descida seguida de subida",
			prices:   []int{6, 5, 4, 7, 8},
			expected: 8, // [6,5,4] = 6 períodos + [7] + [8]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getDescentPeriods(tt.prices)
			if result != tt.expected {
				t.Errorf("getDescentPeriods(%v) = %d; esperado %d", tt.prices, result, tt.expected)
			}
		})
	}
}

func BenchmarkGetDescentPeriods(b *testing.B) {
	prices := make([]int, 1000)
	for i := range prices {
		prices[i] = 1000 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getDescentPeriods(prices)
	}
}

func BenchmarkGetDescentPeriodsWorstCase(b *testing.B) {
	// Pior caso: sequência completamente descendente
	prices := make([]int, 10000)
	for i := range prices {
		prices[i] = 10000 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getDescentPeriods(prices)
	}
}

func BenchmarkGetDescentPeriodsBestCase(b *testing.B) {
	// Melhor caso: sequência crescente (sem períodos descendentes)
	prices := make([]int, 10000)
	for i := range prices {
		prices[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getDescentPeriods(prices)
	}
}
