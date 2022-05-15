package generics

func Sum(numbers []int) int {

	add := func(x, y int) int { return x + y }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	appendSum := func(sums []int, numbers []int) []int {
		return append(sums, Sum(numbers))
	}

	return Reduce(numbersToSum, appendSum, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {

	sumTail := func(sumsTail []int, numbers []int) []int {
		var sum int
		if len(numbers) > 0 {
			sum = Sum(numbers[1:])
		}

		return append(sumsTail, sum)
	}
	return Reduce(numbersToSum, sumTail, []int{})
}
