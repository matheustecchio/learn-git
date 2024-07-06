package main

import "fmt"

var (
	revenue  float64
	expenses float64
	taxRate  float64
)

func getInformations() {
	fmt.Print("Your Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
}

func earningsBeforeTax(revenue, expenses float64) float64 {
	return revenue - expenses
}

func earningsAfterTax(revenue, expenses, taxRate float64) float64 {
	return ((revenue - expenses) * (1 - taxRate/100))
}

func calculateRatio() float64 {
	EBT := earningsBeforeTax(revenue, expenses)
	profit := earningsAfterTax(revenue, expenses, taxRate)

	return EBT / profit
}

func main() {
	getInformations()
	fmt.Printf("\nEarning before tax: %.2f", earningsBeforeTax(revenue, expenses))
	fmt.Printf("\nEarning after tax: %.2f", earningsAfterTax(revenue, expenses, taxRate))
	fmt.Printf("\nRatio: %.2f\n", calculateRatio())

}
