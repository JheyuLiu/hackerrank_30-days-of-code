package main

import "fmt"

func swap_num(a, b *int) {
	*a, *b = *b, *a
}

func bubble_sort(nums []int) int {
	var number_of_swaps int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				swap_num(&nums[j], &nums[j+1])
				number_of_swaps++
			}
		}

		if number_of_swaps == 0 {
			break;
		}
	}

	return number_of_swaps
}

func main() {
	var num int
	fmt.Scan(&num)
    nums := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&nums[i])
	}
	count := bubble_sort(nums)

	fmt.Println("Array is sorted in", count, "swaps.")
	fmt.Println("First Element: ", nums[0])
	fmt.Println("Last Element: ", nums[num-1])
}