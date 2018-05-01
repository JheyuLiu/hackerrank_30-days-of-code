package main

import (
    "fmt"
    "sort"
)

type Diff struct {
	elements []int
	maxDiff int
}

func (d *Diff) ComputeDifference() {
    sort.Ints(d.elements)
    last_index := len(d.elements)-1
    d.maxDiff = d.elements[last_index] - d.elements[0]
}

func main() {
	var num int

    fmt.Scan(&num)
    scores := make([]int, num)
    for i := 0; i < num; i++ {
    	fmt.Scan(&scores[i])
    }

    diff := Diff{elements: scores}
    diff.ComputeDifference()
    fmt.Println(diff.maxDiff)
}