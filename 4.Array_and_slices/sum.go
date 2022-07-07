package arrays

func Sum(numbers []int) int {
	sum := 0

	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	// An alternative version is
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Go let you write variadic functions that can take a variable number of
// arguments
func SumAll(numbersToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersToSum)
	// Create a slice with a starting capacity of the len of the numbers to sum
	//sums := make([]int, lengthOfNumbers)

	//for i, numbers := range numbersToSum {
	//	sums[i] = Sum(numbers)
	//}

	// To prevent accessing positions greater than the slice capacity we
	// can do the following
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

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
