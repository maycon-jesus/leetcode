package main

import (
	"reflect"
	"testing"
)

func TestSortEvents(t *testing.T) {
	tests := []struct {
		name     string
		events   [][]string
		expected [][]string
	}{
		{
			name: "Sort by timestamp ascending",
			events: [][]string{
				{"MESSAGE", "30", "ALL"},
				{"MESSAGE", "10", "HERE"},
				{"MESSAGE", "20", "id1"},
			},
			expected: [][]string{
				{"MESSAGE", "10", "HERE"},
				{"MESSAGE", "20", "id1"},
				{"MESSAGE", "30", "ALL"},
			},
		},
		{
			name: "OFFLINE comes before other events at same timestamp",
			events: [][]string{
				{"MESSAGE", "10", "ALL"},
				{"OFFLINE", "10", "1"},
				{"MESSAGE", "10", "HERE"},
			},
			expected: [][]string{
				{"OFFLINE", "10", "1"},
				{"MESSAGE", "10", "ALL"},
				{"MESSAGE", "10", "HERE"},
			},
		},
		{
			name: "Multiple OFFLINE events at same timestamp",
			events: [][]string{
				{"OFFLINE", "5", "2"},
				{"OFFLINE", "5", "1"},
				{"MESSAGE", "5", "ALL"},
			},
			expected: [][]string{
				{"OFFLINE", "5", "2"},
				{"OFFLINE", "5", "1"},
				{"MESSAGE", "5", "ALL"},
			},
		},
		{
			name:     "Empty events",
			events:   [][]string{},
			expected: [][]string{},
		},
		{
			name: "Single event",
			events: [][]string{
				{"MESSAGE", "10", "ALL"},
			},
			expected: [][]string{
				{"MESSAGE", "10", "ALL"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortEvents(tt.events)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sortEvents() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCountMentions(t *testing.T) {
	tests := []struct {
		name          string
		numberOfUsers int
		events        [][]string
		expected      []int
	}{
		{
			name:          "Example from problem",
			numberOfUsers: 5,
			events: [][]string{
				{"OFFLINE", "28", "1"},
				{"OFFLINE", "14", "2"},
				{"MESSAGE", "2", "ALL"},
				{"MESSAGE", "28", "HERE"},
				{"OFFLINE", "66", "0"},
				{"MESSAGE", "34", "id2"},
				{"MESSAGE", "83", "HERE"},
				{"MESSAGE", "40", "id3 id3 id2 id4 id4"},
			},
			expected: []int{2, 1, 4, 5, 5},
		},
		{
			name:          "MESSAGE ALL - all users get mentioned",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "10", "ALL"},
			},
			expected: []int{1, 1, 1},
		},
		{
			name:          "MESSAGE HERE - only online users",
			numberOfUsers: 3,
			events: [][]string{
				{"OFFLINE", "5", "1"},
				{"MESSAGE", "10", "HERE"},
			},
			expected: []int{1, 0, 1},
		},
		{
			name:          "MESSAGE HERE - all users online",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "10", "HERE"},
			},
			expected: []int{1, 1, 1},
		},
		{
			name:          "Direct mentions - single user",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "10", "id1"},
			},
			expected: []int{0, 1, 0},
		},
		{
			name:          "Direct mentions - multiple users",
			numberOfUsers: 5,
			events: [][]string{
				{"MESSAGE", "10", "id0 id2 id4"},
			},
			expected: []int{1, 0, 1, 0, 1},
		},
		{
			name:          "Direct mentions - repeated mentions",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "10", "id1 id1 id2"},
			},
			expected: []int{0, 2, 1},
		},
		{
			name:          "User goes offline and comes back online after 60 seconds",
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "20", "HERE"}, // User 0 offline, only user 1 mentioned
				{"MESSAGE", "70", "HERE"}, // User 0 back online
			},
			expected: []int{1, 2},
		},
		{
			name:          "User offline exactly at boundary",
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "70", "HERE"}, // Exactly at offline deadline (10+60)
			},
			expected: []int{1, 1},
		},
		{
			name:          "Multiple users offline at different times",
			numberOfUsers: 3,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"OFFLINE", "20", "1"},
				{"MESSAGE", "30", "HERE"}, // Only user 2 is online
				{"MESSAGE", "75", "HERE"}, // All users online again
			},
			expected: []int{1, 0, 2},
		},
		{
			name:          "No events",
			numberOfUsers: 3,
			events:        [][]string{},
			expected:      []int{0, 0, 0},
		},
		{
			name:          "Only OFFLINE events",
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"OFFLINE", "20", "1"},
			},
			expected: []int{0, 0},
		},
		{
			name:          "Mixed events with multiple mentions",
			numberOfUsers: 4,
			events: [][]string{
				{"MESSAGE", "5", "ALL"},
				{"OFFLINE", "10", "1"},
				{"MESSAGE", "15", "id0 id1"},
				{"MESSAGE", "20", "HERE"},
				{"MESSAGE", "80", "id2 id2 id3"},
			},
			expected: []int{3, 2, 4, 3},
		},
		{
			name:          "User goes offline multiple times",
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "20", "HERE"},
				{"OFFLINE", "80", "0"},
				{"MESSAGE", "90", "HERE"},
			},
			expected: []int{0, 2},
		},
		{
			name:          "Single user",
			numberOfUsers: 1,
			events: [][]string{
				{"MESSAGE", "10", "ALL"},
				{"MESSAGE", "20", "id0"},
			},
			expected: []int{2},
		},
		{
			name:          "OFFLINE and MESSAGE at same timestamp - OFFLINE processed first",
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "10", "HERE"},
				{"OFFLINE", "10", "0"},
			},
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countMentions(tt.numberOfUsers, tt.events)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("countMentions() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkCountMentions(b *testing.B) {
	events := [][]string{
		{"OFFLINE", "28", "1"},
		{"OFFLINE", "14", "2"},
		{"MESSAGE", "2", "ALL"},
		{"MESSAGE", "28", "HERE"},
		{"OFFLINE", "66", "0"},
		{"MESSAGE", "34", "id2"},
		{"MESSAGE", "83", "HERE"},
		{"MESSAGE", "40", "id3 id3 id2 id4 id4"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countMentions(5, events)
	}
}

func BenchmarkSortEvents(b *testing.B) {
	events := [][]string{
		{"OFFLINE", "28", "1"},
		{"OFFLINE", "14", "2"},
		{"MESSAGE", "2", "ALL"},
		{"MESSAGE", "28", "HERE"},
		{"OFFLINE", "66", "0"},
		{"MESSAGE", "34", "id2"},
		{"MESSAGE", "83", "HERE"},
		{"MESSAGE", "40", "id3 id3 id2 id4 id4"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortEvents(events)
	}
}
