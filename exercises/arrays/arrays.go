package arrays

// Sum ..
func Sum(numbers []int) int {

	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// using range
	for _, number := range numbers {
		sum += number
	}
	return sum
}
