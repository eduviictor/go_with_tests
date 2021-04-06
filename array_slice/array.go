package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
