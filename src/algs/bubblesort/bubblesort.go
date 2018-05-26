package bubblesort

// BubbleSort sort int array by bubble alg
func BubbleSort(intArr []int) {
	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			if intArr[j] < intArr[i] {
				intArr[i], intArr[j] = intArr[j], intArr[i]
			}
		}
	}
}
