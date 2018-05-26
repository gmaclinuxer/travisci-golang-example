package qsort

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {

	intArr := []int{4, 5, 2, 3, 1, 0}
	fmt.Println("before sort: ", intArr)
	QuickSort(intArr, 0, len(intArr)-1)
	fmt.Println("after sort: ", intArr)
	for index, v := range intArr {
		if index != v {
			t.Errorf("quicksort failed: intaArr[%d] != %d", index, v)
			break
		}
	}
}
