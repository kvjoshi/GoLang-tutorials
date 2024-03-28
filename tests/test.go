package main

import "fmt"

func main() {
	ages := []int{42, 28, 30, 27, 18}

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}

	for i, age := range ages {
		fmt.Println(i, age)
	}

}
