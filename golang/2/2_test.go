package main

import (
	"reflect"
	"testing"
)

func listToSlice(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{[]int{1, 8}, []int{0}, []int{1, 8}},
	}

	for _, tt := range tests {
		l1 := sliceToList(tt.l1)
		l2 := sliceToList(tt.l2)
		result := addTwoNumbers(l1, l2)
		got := listToSlice(result)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; want %v", tt.l1, tt.l2, got, tt.expected)
		}
	}
}
