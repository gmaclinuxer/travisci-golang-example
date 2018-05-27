package bubblesort

/*
bubblesort git:(master) ✗ gotests -all -w ./bubblesort.go
Generated TestBubbleSort
*/

import (
	"testing"
	"reflect"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		intArr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"basic", args{[]int{1, 0, 3, 4, 2}}, []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.intArr)
			if !reflect.DeepEqual(tt.want, tt.args.intArr) {
				t.Error("bubble sort failed")
			}
		})
	}
}
