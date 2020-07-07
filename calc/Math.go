package calc

// Add - Returns sum of given 2 integers
func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
