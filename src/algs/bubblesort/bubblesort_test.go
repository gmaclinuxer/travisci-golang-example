package bubblesort

import (
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	intArr := []int{4,5,2,3,1,0}
	fmt.Println("before sort: ", intArr)
	BubbleSort(intArr)
	fmt.Println("after sort: ", intArr)
	for index, v := range intArr {
		if index != v {
			t.Errorf("bubblesort failed: intaArr[%d] != %d", index, v)
			break
		}
	}
}