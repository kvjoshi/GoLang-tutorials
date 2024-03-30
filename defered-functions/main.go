package main

import (
	"fmt"
)

type Engineer struct {
	Name string
}

func (e *Engineer) UpdateName(name string) {
	e.Name = name
}
func doStuff(e *Engineer) {
	// this will be considered in the defered function
	name := "Karmavir Joshi"
	defer e.UpdateName(name)
	// this will not be used as the defered function will be called after this but defination was before the deffered function
	name = "Karmavir N Joshi"

}

func main() {
	fmt.Println("main")

	kv := &Engineer{Name: "Karmavir"}
	fmt.Printf("%+v\n", kv)
	doStuff(kv)
	fmt.Printf("%+v\n", kv)
}
