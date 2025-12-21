package chapter3

import (
	"fmt"
)

func E1() {

}

func PrintSlice(x []int) {
	fmt.Println(x, len(x), cap(x))
}
