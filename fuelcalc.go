package main

import (
	"fmt"

	"github.com/lexbarker/fuelcalc"
)

//import "fmt"

func main() {
	var fc fuelcalc.Mpg
	var fu fuelcalc.Fuelused
	fmt.Printf("How many Miles?\n")
	fmt.Scan(&fc.Miles)
	fmt.Printf("\nHow many Litres ?\n")
	fmt.Scan(&fc.Litres)
	res := fc.Consumption()
	fmt.Printf("MPG returned is %f\n", res)
	fmt.Printf("How many Miles ? \n")
	fmt.Scan(&fu.Miles)
	fmt.Printf("\nwhat is the MPG?\n")
	fmt.Scan(&fu.Mpg)
	fu.Fuelrequired()
	fmt.Printf("\nyou need %f litres", fu.Litres)
}
