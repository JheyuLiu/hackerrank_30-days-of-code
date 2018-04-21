package main

import (
	"fmt"
	"math"
)

func main() {
	var cost float64
	var tip, tax int

	fmt.Scan(&cost)
	fmt.Scan(&tip)
	fmt.Scan(&tax)

	cost += (cost * (float64(tip) / 100)) + (cost * (float64(tax) / 100))
	fmt.Printf("The total meal cost is %d dollars.\n", int(math.Floor(cost+0.5)))
}
