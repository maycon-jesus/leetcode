package main

import "testing"

func Test_partitionArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{3, 6, 1, 2, 5},
				k:    2,
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 2, 3},
				k:    1,
			},
			want: 2,
		},
		{
			name: "test 3",
			args: args{
				nums: []int{2, 2, 4, 5},
				k:    0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("partitionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_partitionArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitionArray([]int{3, 6, 1, 2, 5}, 2)
	}
}
