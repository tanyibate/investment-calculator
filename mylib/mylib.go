package mylib

import "math"

func CalculateReturn(startingValue int, interestRate float64, numOfYears int) float64 {

	var futureValue float64 = float64(startingValue) * math.Pow(1+(interestRate/100), float64(numOfYears))

	return futureValue
}
