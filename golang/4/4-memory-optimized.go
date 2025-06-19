package main

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	totalSize := len(nums1) + len(nums2)
// 	if totalSize == 0 {
// 		return 0
// 	}

// 	var m1, m2 int
// 	i, j := 0, 0

// 	for k := 0; k <= totalSize/2; k++ {
// 		m1 = m2
// 		if i < len(nums1) && (j >= len(nums2) || nums1[i] < nums2[j]) {
// 			m2 = nums1[i]
// 			i++
// 		} else {
// 			m2 = nums2[j]
// 			j++
// 		}
// 	}

// 	if totalSize%2 == 0 {
// 		return float64(m1+m2) / 2.0
// 	}

// 	return float64(m2)
// }
