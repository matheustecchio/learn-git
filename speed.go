package main

import "fmt"

func t() {
	i := 0

	for i < 10000000 {
		fmt.Print(i)
		i++
	}
}
