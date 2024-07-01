package main

import "fmt"

func main() {
	// Slices: Can increase and decrease in capacity. As the capacity reach its top limit, Go double its capacity in memory.
	// The first capacity of a slice declared with only `:=` is the length of it.
	names := []string{"Tiago", "Daniel", "Marcos", "Marta"}
	fmt.Println(names, len(names), cap(names))

	names = append(names, "Rafael")
	fmt.Println(names, len(names), cap(names))

	names = append(names, "Roberta")
	fmt.Println(names, len(names), cap(names))

	names = append(names, "Rafael")
	fmt.Println(names, len(names), cap(names))

	names = append(names, "Matheus")
	fmt.Println(names, len(names), cap(names))

	names = append(names, "John")
	fmt.Println(names, len(names), cap(names))

	// If we want to make slice with a capacity bigger than the length of it, we can use the make() function.
	students := make([]string, 0, 15) // It starts the slice with 0 values but with 15 of capacity in memory.
	fmt.Println(students, len(students), cap(students))

	pizza := make([]string, 8) // It starts the slice with 8 values and a capacity of 8 as well.
	fmt.Println(pizza, len(pizza), cap(pizza))

}
