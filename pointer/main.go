package main

import "fmt"

type Engineer struct {
	Name string
	Age  int
}

func (e *Engineer) UpdateAge() {
	e.Age += 1
}
func (e *Engineer) UpdateName() {
	e.Name = "Karmavir Joshi"

}

func main() {

	karmavir := &Engineer{
		Name: "Karmavir",
		Age:  27,
	}
	fmt.Println(karmavir)
	karmavir.UpdateAge()
	fmt.Println(karmavir)
	karmavir.UpdateName()
	fmt.Println(karmavir)
}
