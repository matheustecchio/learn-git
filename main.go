package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.0

	var (
		investmentAmount   float64
		expectedReturnRate float64
		years              float64
	)

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	var realFutureValue = futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Real Future Value: %.2f\n", realFutureValue)

}
