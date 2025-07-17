package main

import (
	"slices"
)

func mostBooked(n int, meetings [][]int) int {
	rooms := make([]int, n)
	booked := make([]int, n)

	slices.SortFunc(meetings, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		assigned := false
		for i := 0; i < n; i++ {
			if rooms[i] <= start {
				rooms[i] = end
				booked[i]++
				assigned = true
				break
			}
		}
		if !assigned {
			minEnd := rooms[0]
			minIdx := 0
			for i := 1; i < n; i++ {
				if rooms[i] < minEnd {
					minEnd = rooms[i]
					minIdx = i
				}
			}
			rooms[minIdx] += end - start
			booked[minIdx]++
		}
	}

	maxBooked := booked[0]
	ans := 0
	for i := 1; i < n; i++ {
		if booked[i] > maxBooked {
			maxBooked = booked[i]
			ans = i
		}
	}
	return ans
}
