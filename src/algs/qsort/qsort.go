package qsort

func QuickSort(intArr []int, left, right int) {
	if left < right {

		key := intArr[(left + right/2)]
		i, j := left, right

		for {

			// walk forward larger item
			for intArr[i] < key {
				i++
			}

			// walk backward smaller item
			for intArr[j] > key {
				j--
			}

			// walk finished
			if i >= j {
				break
			}

			// swap i and j
			intArr[i], intArr[j] = intArr[j], intArr[i]

		}

		QuickSort(intArr, left, i-1)
		QuickSort(intArr, j+1, right)
	}
}
