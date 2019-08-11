package main

import "fmt"

// Sum ..
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll ... variadic
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// append function which takes a slice and a new value, returning a new slice with all the items in it.
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails ...
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	n := []int{1, 9, 5}

	fmt.Println(Sum(n))
}
