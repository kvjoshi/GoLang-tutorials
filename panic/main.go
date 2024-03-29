package main

import "fmt"

func IPanicEasily() {
	panic("I'm panicking!")
}

func main() {
	fmt.Println("Panic! in the Go App")
	IPanicEasily()
}
