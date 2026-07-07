package iteration

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {

	var sums []int

	for _, numslice := range numbers {
		if len(numslice) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numslice[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
