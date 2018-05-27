package qsort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	// table driven tests
	tests := []struct {
		arg  []int
		want []int
	}{
		{arg: []int{4, 5, 2, 3, 1, 0}, want: []int{0, 1, 2, 3, 4, 5}},
	}

	for _, tc := range tests {
		intArr := tc.arg
		wantArr := tc.want

		fmt.Println("before sort: ", intArr)
		QuickSort(intArr, 0, len(intArr)-1)
		fmt.Println("after sort: ", intArr)
		for index, v := range intArr {
			if wantArr[index] != v {
				t.Errorf("quicksort failed: intaArr[%d] != %d", index, v)
				break
			}
		}
	}
}
