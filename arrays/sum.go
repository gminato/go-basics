package main

func Sum(numbers []int) int {
	sum := 0
	for _, ele := range numbers {
		sum += ele
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
