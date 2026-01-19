package chapter5

import "fmt"

func E3(prefix string) func(input string) {
	return func(input string) {
		fmt.Println(prefix, input)
	}
}
