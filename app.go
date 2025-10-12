package main //main tells go that main package is the main entry point
import (
	"fmt"
	"math"
)

// uint - non negative integer data type, int=0, string="", float64=0.0, bool= false, int32
func main() {
	const inflationRate = 2.5
	var years float64 = 10
	var investmentAmount float64
	expectedReturnRate := 5.5 //short notation to declare a variable

	//input
	fmt.Print("Please enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//output
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
