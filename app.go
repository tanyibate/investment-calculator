package main

import (
	"example.com/investment-calculator/mylib"
	"fmt"
)

func main() {
	var years, investmentAmount int
	var interestRate float64 = 4.5
	const inflationRate = 6.5

	fmt.Print("Whats the investment amount?:")
	fmt.Scan(&investmentAmount)

	fmt.Println("For how many years?:")
	fmt.Scan(&years)

	var futureValue float64 = mylib.CalculateReturn(investmentAmount, interestRate, years)
	var futureRealValue = futureValue / (1 + (inflationRate / 100))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
