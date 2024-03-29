package main

import (
	"errors"
	"fmt"
)

func IPanicEasily() {
	panic("I'm panicking!")
}

func MyFunction() (err error) {
	fmt.Println("Starting my function")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in MyFunction", r)
			err = errors.New("I panicked easily in MyFunction")
		}
	}()
	IPanicEasily()
	fmt.Println("Ending my function")
	return nil
}

func main() {
	fmt.Println("Panic! in the Go App")
	if err := MyFunction(); err != nil {
		fmt.Println("Error in MyFunction", err)
		fmt.Println(err.Error())
	}
}
