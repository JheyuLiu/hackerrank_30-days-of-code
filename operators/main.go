package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cost, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	tip, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	scanner.Scan()
	tax, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	cost += (cost * (float64(tip) / 100)) + (cost * (float64(tax) / 100))
	fmt.Printf("The total meal cost is %d dollars.\n", int(math.Floor(cost+0.5)))
}
