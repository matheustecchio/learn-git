package main

import "fmt"

func main() {
	// Classic loop: start; conditional; after
	for i := 0; i < 10; i = i + 2 {
		fmt.Println(i)
	} // You could put `start` before the for loop, and the `after` just inside the for loop as well, it's up to you.

	names := []string{"Tiago", "Daniel", "Maria", "Marta"}

	// Looping with the length of the array.
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], i)
	}
	// For and Range
	for index, name := range names {
		fmt.Println(index, name)
	} // You can also use `_` instead of index, so Go will ignore it and just get the name data for you.

	// Infinite loop
	for {
		fmt.Println("loop")
		break
	}

}
