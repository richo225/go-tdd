package sum

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numberArrays ...[]int) []int {
	sums := make([]int, len(numberArrays))

	for i, numberArray := range numberArrays {
		sums[i] = Sum(numberArray)
	}

	return sums
}
