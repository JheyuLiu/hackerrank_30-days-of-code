package main

import "fmt"

type AdvancedArithmetic interface {
	divisorSum(num int) int
}

type Calculator struct {

}

func (cal Calculator) divisorSum(num int) int {
    var sum int
    for i := 1; i <= num; i++ {
    	if num % i == 0 {
    		sum += i
    	}
    }
    return sum
}

func check(something interface{}) bool {
    if _, ok := something.(AdvancedArithmetic); ok {
    	return ok;
    } 
    return false;
}

func main() {
	var num int
	fmt.Scan(&num)

	cal := Calculator{}
	if check(cal) {
		sum := cal.divisorSum(num)
		fmt.Println("I implemented: AdvancedArithmetic");
		fmt.Println(sum)
	} else {
		fmt.Println("Wrong answer");
	}
}