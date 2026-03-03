package bubblesort

import "fmt"

func BubbleSort(unsorted []int) []int {
	for i, v := range unsorted {
		fmt.Println("index: ", i, " value: ", v)
	}
	return unsorted
}
