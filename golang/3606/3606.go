package main

import (
	"slices"
	"strings"
)

var validBusinessLines = map[string]bool{
	"grocery":     true,
	"electronics": true,
	"pharmacy":    true,
	"restaurant":  true,
}

var businessOrder = map[string]int{
	"electronics": 0,
	"grocery":     1,
	"pharmacy":    2,
	"restaurant":  3,
}

func isValidCouponCode(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '_') {
			return false
		}
	}
	return true
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	validIndices := make([]int, 0, len(code))

	for i, coupon := range code {
		size := len(coupon)
		active := isActive[i]
		_, business := validBusinessLines[businessLine[i]]
		if !active || !business || size == 0 {
			continue
		}

		if isValidCouponCode(coupon) {
			validIndices = append(validIndices, i)
		}
	}

	// Ordena usando índices (sem duplicar strings)
	slices.SortFunc(validIndices, func(a, b int) int {
		// Compare business lines
		orderA := businessOrder[businessLine[a]]
		orderB := businessOrder[businessLine[b]]
		if orderA != orderB {
			return orderA - orderB
		}
		// If same business line, compare codes lexicographically
		return strings.Compare(code[a], code[b])
	})

	// Constrói resultado apenas uma vez
	result := make([]string, len(validIndices))
	for i, idx := range validIndices {
		result[i] = code[idx]
	}

	return result
}
