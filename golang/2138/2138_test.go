package main

import (
	"reflect"
	"testing"
)

func TestDivideString(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		fill string
		want []string
	}{
		{
			s:    "abcdefghi",
			k:    3,
			fill: "x",
			want: []string{"abc", "def", "ghi"},
		},
		{
			s:    "abcdefghij",
			k:    3,
			fill: "x",
			want: []string{"abc", "def", "ghi", "jxx"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := divideString(tt.s, tt.k, tt.fill[0])
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideString(%q, %d, %q) = %v; want %v", tt.s, tt.k, tt.fill, got, tt.want)
			}
		})
	}
}
