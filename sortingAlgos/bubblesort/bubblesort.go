package bubblesort

func BubbleSort(unsorted []int) []int {
	for j := 0; j < len(unsorted)-1; j++ {
		for i := 0; i < len(unsorted)-1-j; i++ {
			if unsorted[i] > unsorted[i+1] {
				temp := unsorted[i+1]
				unsorted[i+1] = unsorted[i]
				unsorted[i] = temp
			}
		}
	}
	return unsorted
}
