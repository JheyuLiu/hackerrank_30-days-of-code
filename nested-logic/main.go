package main

import "fmt"

func main() {
	var actualDay, actualMonth, actualYear int
	var expectDay, expectMonth, expectYear int

	fmt.Scan(&actualDay, &actualMonth, &actualYear)
	fmt.Scan(&expectDay, &expectMonth, &expectYear)

	var fine int
	if actualYear > expectYear {
		fine = 10000
	} else if actualMonth > expectMonth && actualYear == expectYear {
		fine = fine + 500 * (actualMonth - expectMonth)
	} else if actualDay > expectDay && actualMonth == expectMonth && actualYear == expectYear  {
		fine = fine + 15 * (actualDay - expectDay)
	} else {
		fine = 0
	}

	fmt.Println(fine)
}