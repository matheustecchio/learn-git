package main

import (
	"fmt"
	"log"
	"os"
)

func readFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}

var heads, tails int

func flipCoin(side string) {

	switch side {
	case "head":
		heads++
	case "tail":
		tails++
	default:
		fmt.Println("Coin is straight")
	}

	fmt.Print("\nHeads: ", heads, "\nTails: ", tails, "\n")
}

func main() {
	readFile("hello.txt")

	flipCoin("head")
	flipCoin("head")
	flipCoin("tail")

}
