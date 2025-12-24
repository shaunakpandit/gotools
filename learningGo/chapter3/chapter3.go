package chapter3

import (
	"fmt"
)

func All() {
	E1()
	E2()
	E3()
}

func E1() {
	greetings := []string{"hello", "Hola", "Hindi", "Chinese", "russian"}
	g1 := greetings[:2]
	g2 := greetings[1:4]
	g3 := greetings[3:]

	fmt.Println(greetings, g1, g2, g3)
}

func E2() {
	message := "Hi 🤦‍♂️ and 😂"
	var rs []rune = []rune(message)
	fmt.Println(string(rs[3]))
}

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func E3() {
	emp1 := Employee{"john", "smith", 0}
	emp2 := Employee{
		firstName: "Mark",
		lastName:  "Anthony",
		id:        1}
	var emp3 Employee
	emp3.firstName = "hi"
	emp3.lastName = "there"
	emp3.id = 2

	fmt.Println(emp1, emp2, emp3)
}

func PrintSlice(x []int) {
	fmt.Println(x, len(x), cap(x))
}
