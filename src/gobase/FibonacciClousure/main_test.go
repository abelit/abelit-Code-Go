package main

import (
	"reflect"
	"testing"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "fib test",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib()(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
