package syracuse

func Syracuse(n int) []int {
	numbers := make([]int, 0)
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		numbers = append(numbers, n)
	}

	return numbers
}
