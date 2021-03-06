// Problem 6 - http://projecteuler.net/
//
// Find the difference between the sum of the squares of the
// first one hundred natural numbers and the square of the sum.

package projecteuler

import (
	"fmt"
	"math"
)

func square_of_sum(iterations int) int {
	sum := 0
	for i := 1; i <= iterations; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func sum_of_squares(iterations float64) int {
	sum := float64(0)
	for i := float64(1); i <= iterations; i++ {
		sum += math.Pow(i, 2)
	}
	return int(sum)
}

func Euler6() string {
	diff := square_of_sum(100) - sum_of_squares(100)
	return fmt.Sprint(diff)
}
