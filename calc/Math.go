package calc

import "errors"

// Add - Returns sum of given 2 integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("Should give more than 2 numbers")
	}
	for _, num := range numbers {
		sum = sum + num
	}
	return sum, nil
}
