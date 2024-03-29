package main

import "fmt"

func main() {
	var name string
	name = "Elliot"
	fmt.Println(name)
	ptr := &name

	// The value of the pointer is the memory address of the variable.
	fmt.Println(ptr)
	// The value of the pointer is the value of the variable.
	fmt.Println(*ptr)

	// Change the value of the variable.
	*ptr = "Forrest"
	fmt.Println(name)

}
