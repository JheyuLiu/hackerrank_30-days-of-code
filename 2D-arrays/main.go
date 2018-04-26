package main

import "fmt"

const (
	ROW = 6
	COL = 6
)

func main() {
	arr := make([][]int, ROW)
	for i := range arr {
		arr[i] = make([]int, COL)
	}

	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

    var max, sum int = -100, 0
    for i := 0; i < ROW-2; i++ {
		for j := 0; j < COL-2; j++ {
			sum = arr[i][j]
			sum += arr[i][j+1]
			sum += arr[i][j+2]
            sum += arr[i+1][j+1]
            sum += arr[i+2][j]
            sum += arr[i+2][j+1]
            sum += arr[i+2][j+2]

			if sum > max {
				max = sum
			}
		}
	}	
	fmt.Println(max)
}