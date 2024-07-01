package main

import "fmt"

func main() {
	ages := make(map[string]int8)
	ages["Tiago"] = 31
	ages["Marta"] = 43
	ages["Daniel"] = 26
	ages["Matheus"] = 23

	fmt.Println(ages)
	fmt.Println(ages["Tiago"])
	fmt.Println(ages["Lucas"]) // As `Lucas` doesn't exist, it returns the zero-value of the type, in this case, 0.

	// We can assign two variables to a map value, the first return the value itself and the second returns a boolean, true if the key exists and false if not.
	value, ok := ages["Marta"]
	fmt.Println(value, ok)

}
