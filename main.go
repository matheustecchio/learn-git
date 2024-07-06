package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var balance float64

func displayMenu() {
	fmt.Println("\nWelcome to the Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("0. Exit")
	var choice uint8
	fmt.Print("Enter with you option: ")

	if _, err := fmt.Scan(&choice); err != nil {
		log.Panic("Error while computing menu's choice")
	}

	switch choice {
	case 1:
		checkBalance()
	case 2:
		depositMoney()
	case 3:
		withdrawMoney()
	case 0:
		os.Exit(0)
	}

}

func checkBalance() {
	fmt.Printf("Your balance is: %.2f\n", balance)
}

func depositMoney() {
	fmt.Println("Insert how much do you wan to deposit: ")
	var amount float64
	if _, err := fmt.Scan(&amount); err != nil {
		log.Fatal("\nError while depositing...")
	}

	balance = balance + amount

	fmt.Printf("\nYou deposited %2.f to your bank account!\n", amount)
}

func withdrawMoney() {
	fmt.Println("Insert how much do you wan to withdraw: ")
	var amount float64
	if _, err := fmt.Scan(&amount); err != nil {
		log.Fatal("\nError while withdrawing...")
	}

	if amount > balance {
		log.Fatal("\nYou don't have enough money to complete this operation!")
	}

	balance = balance - amount

	fmt.Printf("\nYou withdrew %2.f to your bank account!\n", amount)

}

// If a file doesn't exist, it creates the file with the designated name and format, and insert the value 0.
func createFile(fileName string) {
	_, err := os.ReadFile(fileName)
	if err != nil {
		os.Create(fileName)
		os.WriteFile(fileName, []byte("0"), 0644)
	}
}

// Synchronizes the file data to a variable in our code.
func syncData(fileName string) float64 {
	fileByte, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("File '%s' doesn't exist or cannot be read: %v", fileName, err)
	}

	byteToString := string(fileByte)
	stringToFloat64, err := strconv.ParseFloat(byteToString, 64)
	if err != nil {
		log.Fatalf("Failed to convert balance data from file '%s' to float64: %v", fileName, err)
	}

	return stringToFloat64
}

func saveDataToFile(fileName string, targetFloat float64) {
	floatToString := fmt.Sprint(targetFloat)
	os.WriteFile(fileName, []byte(floatToString), 0644)
}

func main() {
	createFile("balance.txt")
	balance = syncData("balance.txt")

	for {
		displayMenu()
		saveDataToFile("balance.txt", balance)
	}

}
