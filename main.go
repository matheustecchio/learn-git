package main

import (
	"fmt"
	"math"
	"strconv"
)

func greet(name ...string) {
	var userName string
	if len(name) > 0 {
		userName = name[0]
	} else {
		userName = "User"
	}
	fmt.Println("Hi", userName)
}

func convertAndSum(a int, b string) (int, error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}

	total := a + i

	return total, nil
}

func main() {
	fmt.Println(math.Pi)
	name := "matheus"
	fmt.Println(name)

	greet("Matheus")

	fmt.Println(convertAndSum(10, "5"))
}
