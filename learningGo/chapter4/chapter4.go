package chapter4

import "fmt"

func All() {
	E1()
	E2()
	E3()
}

func E1() []int {
	slice := CreateSlice()
	fmt.Println(slice)
	return slice
}

func E2() {
	slice := CreateSlice()
loop:
	for _, v := range slice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
			break loop
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func E3() {
	var total int
	for i := range 10 {
		total := total + i
		fmt.Println(total)
	}

	// prints as 0 as it is a shadow var in the loop!
	fmt.Println(total)
}

func CreateSlice() []int {
	slice := []int{}
	for i := range 100 {
		slice = append(slice, i+1)
	}
	return slice
}
