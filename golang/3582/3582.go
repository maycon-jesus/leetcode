package main

import (
	"strings"
)

// Problem: Generate Tag from Video Caption
// Link: https://leetcode.com/problems/generate-tag-from-video-caption/
// Dificulty: Easy
// Tags: String, Simulation
// Time Complexity: O(n), where n is the length of the caption
// Space Complexity: O(n), where n is the length of the tag generated

func generateTag(caption string) string {
	strBuilder := strings.Builder{}
	capitalizeIndice := -1
	letterFound := false
	for i, char := range caption {
		capitalizeNext := i == capitalizeIndice

		if char != ' ' {
			letterFound = true // Found a letter, so we can start building the tag
		}

		if char == ' ' && letterFound {
			capitalizeIndice = i + 1 // Next character should be capitalized
		} else if !capitalizeNext && char >= 'A' && char <= 'Z' {
			strBuilder.WriteRune(char + 32) // Convert uppercase to lowercase
		} else if capitalizeNext && char >= 'a' && char <= 'z' {
			strBuilder.WriteRune(char - 32) // Keep uppercase as is
		} else if letterFound {
			strBuilder.WriteRune(char)
		}
		if strBuilder.Len() == 99 {
			break
		}
	}
	return "#" + strBuilder.String()
}
