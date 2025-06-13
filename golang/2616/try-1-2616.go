package main

// import (
// 	"fmt"
// 	"slices"
// )

// func quickSort(nums []int, left, right int) {
// 	if left >= right {
// 		return
// 	}
// 	pivot := nums[(left+right)/2]
// 	i, j := left, right
// 	for i <= j {
// 		for nums[i] < pivot {
// 			i++
// 		}
// 		for nums[j] > pivot {
// 			j--
// 		}
// 		if i <= j {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			i++
// 			j--
// 		}
// 	}
// 	if left < j {
// 		quickSort(nums, left, j)
// 	}
// 	if i < right {
// 		quickSort(nums, i, right)
// 	}
// }

// func reverseQuickSort(nums []int, left, right int) {
// 	if left >= right {
// 		return
// 	}
// 	pivot := nums[(left+right)/2]
// 	i, j := left, right
// 	for i <= j {
// 		for nums[i] > pivot {
// 			i++
// 		}
// 		for nums[j] < pivot {
// 			j--
// 		}
// 		if i <= j {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			i++
// 			j--
// 		}
// 	}
// 	if left < j {
// 		reverseQuickSort(nums, left, j)
// 	}
// 	if i < right {
// 		reverseQuickSort(nums, i, right)
// 	}
// }

// func minimizeMax(nums []int, p int) int {
// 	if p == 0 {
// 		return 0
// 	}

// 	reverseQuickSort(nums, 0, len(nums)-1)
// 	solved := false
// 	pairs := make([]int, 0)
// 	ignoredIndex := make(map[int]bool)

// 	for target := 0; !solved; target++ {
// 		for i1, num1 := range nums {
// 			if _, ok := ignoredIndex[i1]; ok {
// 				continue
// 			}
// 			for i2, num2 := range nums[i1+1:] {
// 				i2Real := i2 + i1 + 1
// 				if _, ok := ignoredIndex[i2Real]; ok {
// 					continue
// 				}
// 				if num1-num2 == target {
// 					pairs = append(pairs, num1-num2)
// 					ignoredIndex[i1] = true
// 					ignoredIndex[i2Real] = true
// 					fmt.Println(nums, i1, i2Real)
// 					if p == len(pairs) {
// 						solved = true
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return slices.Max(pairs) // Return the maximum difference found in pairs
// }
