package main

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "test 2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "test 3",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0.0,
		},
		{
			name: "test 4",
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1.0,
		},
		{
			name: "test 5",
			args: args{
				nums1: []int{2},
				nums2: []int{},
			},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	nums1 := make([]int, 1000)
	nums2 := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums1[i] = i
		nums2[i] = i + 1000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
}
