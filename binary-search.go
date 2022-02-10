package main

import "fmt"

func main() {
	sl := []int{11, 22, 33, 44, 55, 66, 77, 88}
	fmt.Println(binarySearch(sl, 88))
}

func binarySearch(sl []int, x int) bool {

	initial := 0
	max := len(sl) - 1

	for initial < max {
		median := (initial + max) / 2
		if x <= sl[median] {
			max = median
		} else {
			initial = median + 1
		}
	}

	if sl[initial] == x || sl[max] == x {
		return true
	}
	return false
}
