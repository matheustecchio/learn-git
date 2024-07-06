package main

import (
	"fmt"
	"log"
)

// FinancialData holds the financial metrics.
type FinancialData struct {
	revenue  float64
	expenses float64
	taxRate  float64
}

// getInput collects user inputs for financial data.
func getInput() FinancialData {
	var data FinancialData

	fmt.Print("Enter your Revenue: ")
	if _, err := fmt.Scan(&data.revenue); err != nil {
		log.Fatalf("Failed to read revenue: %v", err)
	}

	fmt.Print("Enter your Expenses: ")
	if _, err := fmt.Scan(&data.expenses); err != nil {
		log.Fatalf("Failed to read expenses: %v", err)
	}

	fmt.Print("Enter your Tax Rate: ")
	if _, err := fmt.Scan(&data.taxRate); err != nil {
		log.Fatalf("Failed to read tax rate: %v", err)
	}

	return data
}

// earningsBeforeTax calculates earnings before tax.
func earningsBeforeTax(revenue, expenses float64) float64 {
	return revenue - expenses
}

// earningsAfterTax calculates earnings after tax.
func earningsAfterTax(revenue, expenses, taxRate float64) float64 {
	return (revenue - expenses) * (1 - taxRate/100)
}

// calculateRatio calculates the ratio of EBT to profit.
func calculateRatio(data FinancialData) float64 {
	EBT := earningsBeforeTax(data.revenue, data.expenses)
	profit := earningsAfterTax(data.revenue, data.expenses, data.taxRate)

	if profit == 0 {
		log.Fatalf("Profit after tax is zero, cannot calculate ratio")
	}

	return EBT / profit
}

func main() {
	data := getInput()

	EBT := earningsBeforeTax(data.revenue, data.expenses)
	profit := earningsAfterTax(data.revenue, data.expenses, data.taxRate)
	ratio := calculateRatio(data)

	fmt.Printf("\nEarnings before tax: %.2f", EBT)
	fmt.Printf("\nEarnings after tax: %.2f", profit)
	fmt.Printf("\nRatio: %.2f\n", ratio)
}
