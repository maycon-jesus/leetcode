package main

import (
	"reflect"
	"testing"
)

func TestIsValidCouponCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid alphanumeric with underscore",
			input:    "ABC_123",
			expected: true,
		},
		{
			name:     "valid lowercase",
			input:    "abc",
			expected: true,
		},
		{
			name:     "valid uppercase",
			input:    "ABC",
			expected: true,
		},
		{
			name:     "valid numbers",
			input:    "123",
			expected: true,
		},
		{
			name:     "valid underscore only",
			input:    "___",
			expected: true,
		},
		{
			name:     "valid mixed",
			input:    "aB_1cD_2",
			expected: true,
		},
		{
			name:     "invalid with hyphen",
			input:    "ABC-123",
			expected: false,
		},
		{
			name:     "invalid with space",
			input:    "ABC 123",
			expected: false,
		},
		{
			name:     "invalid with special char",
			input:    "ABC@123",
			expected: false,
		},
		{
			name:     "invalid with dot",
			input:    "ABC.123",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidCouponCode(tt.input)
			if result != tt.expected {
				t.Errorf("isValidCouponCode(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidateCoupons(t *testing.T) {
	tests := []struct {
		name         string
		code         []string
		businessLine []string
		isActive     []bool
		expected     []string
	}{
		{
			name:         "basic valid coupons",
			code:         []string{"CODE1", "CODE2", "CODE3"},
			businessLine: []string{"grocery", "electronics", "pharmacy"},
			isActive:     []bool{true, true, true},
			expected:     []string{"CODE2", "CODE1", "CODE3"},
		},
		{
			name:         "filter inactive coupons",
			code:         []string{"CODE1", "CODE2", "CODE3"},
			businessLine: []string{"grocery", "electronics", "pharmacy"},
			isActive:     []bool{true, false, true},
			expected:     []string{"CODE1", "CODE3"},
		},
		{
			name:         "filter invalid business lines",
			code:         []string{"CODE1", "CODE2", "CODE3"},
			businessLine: []string{"grocery", "invalid", "pharmacy"},
			isActive:     []bool{true, true, true},
			expected:     []string{"CODE1", "CODE3"},
		},
		{
			name:         "filter invalid coupon codes",
			code:         []string{"CODE1", "CODE-2", "CODE3"},
			businessLine: []string{"grocery", "electronics", "pharmacy"},
			isActive:     []bool{true, true, true},
			expected:     []string{"CODE1", "CODE3"},
		},
		{
			name:         "filter empty coupon codes",
			code:         []string{"CODE1", "", "CODE3"},
			businessLine: []string{"grocery", "electronics", "pharmacy"},
			isActive:     []bool{true, true, true},
			expected:     []string{"CODE1", "CODE3"},
		},
		{
			name:         "sort by business order",
			code:         []string{"REST1", "ELEC1", "GROC1", "PHAR1"},
			businessLine: []string{"restaurant", "electronics", "grocery", "pharmacy"},
			isActive:     []bool{true, true, true, true},
			expected:     []string{"ELEC1", "GROC1", "PHAR1", "REST1"},
		},
		{
			name:         "sort lexicographically within same business",
			code:         []string{"GROC_C", "GROC_A", "GROC_B"},
			businessLine: []string{"grocery", "grocery", "grocery"},
			isActive:     []bool{true, true, true},
			expected:     []string{"GROC_A", "GROC_B", "GROC_C"},
		},
		{
			name:         "mixed sorting: business then lexicographic",
			code:         []string{"GROC_B", "ELEC_A", "GROC_A", "ELEC_B"},
			businessLine: []string{"grocery", "electronics", "grocery", "electronics"},
			isActive:     []bool{true, true, true, true},
			expected:     []string{"ELEC_A", "ELEC_B", "GROC_A", "GROC_B"},
		},
		{
			name:         "no valid coupons",
			code:         []string{"CODE-1", "CODE-2"},
			businessLine: []string{"grocery", "electronics"},
			isActive:     []bool{true, true},
			expected:     []string{},
		},
		{
			name:         "all filters combined",
			code:         []string{"VALID1", "INVALID-", "", "VALID2", "VALID3"},
			businessLine: []string{"grocery", "electronics", "pharmacy", "invalid", "restaurant"},
			isActive:     []bool{true, true, true, true, false},
			expected:     []string{"VALID1"},
		},
		{
			name:         "empty input",
			code:         []string{},
			businessLine: []string{},
			isActive:     []bool{},
			expected:     []string{},
		},
		{
			name:         "special characters in code",
			code:         []string{"ABC@123", "ABC_123", "ABC 123"},
			businessLine: []string{"grocery", "grocery", "grocery"},
			isActive:     []bool{true, true, true},
			expected:     []string{"ABC_123"},
		},
		{
			name:         "case sensitive sorting",
			code:         []string{"abc", "ABC", "Abc"},
			businessLine: []string{"grocery", "grocery", "grocery"},
			isActive:     []bool{true, true, true},
			expected:     []string{"ABC", "Abc", "abc"},
		},
		{
			name:         "all business lines",
			code:         []string{"R1", "P1", "G1", "E1"},
			businessLine: []string{"restaurant", "pharmacy", "grocery", "electronics"},
			isActive:     []bool{true, true, true, true},
			expected:     []string{"E1", "G1", "P1", "R1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateCoupons(tt.code, tt.businessLine, tt.isActive)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("validateCoupons() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkIsValidCouponCode(b *testing.B) {
	testCases := []string{
		"ABC_123",
		"invalid-code",
		"ValidCode_123",
		"",
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isValidCouponCode(tc)
			}
		})
	}
}

func BenchmarkValidateCoupons(b *testing.B) {
	code := []string{"CODE1", "CODE2", "CODE3", "CODE4", "CODE5"}
	businessLine := []string{"grocery", "electronics", "pharmacy", "restaurant", "grocery"}
	isActive := []bool{true, true, true, true, true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateCoupons(code, businessLine, isActive)
	}
}

func BenchmarkValidateCouponsLarge(b *testing.B) {
	size := 1000
	code := make([]string, size)
	businessLine := make([]string, size)
	isActive := make([]bool, size)
	businesses := []string{"grocery", "electronics", "pharmacy", "restaurant"}

	for i := 0; i < size; i++ {
		code[i] = "CODE_" + string(rune('A'+i%26))
		businessLine[i] = businesses[i%4]
		isActive[i] = i%2 == 0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validateCoupons(code, businessLine, isActive)
	}
}
