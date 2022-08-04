package sum

func Sum(numbers [3]int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}
