//Problem 30
//Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

package projecteuler

import (
	"fmt"
	"strconv"
)

func pow(n int) int {
	return n * n * n * n * n
}

func Euler30() string {
	sum := 0
	for i := 10; i < 1000000; i++ {
		digits := strconv.Itoa(i)
		p := 0
		for _, digit := range digits {
			a, _ := strconv.Atoi(string(digit))
			p += pow(a)
		}
		if i == p {
			sum += i
		}
	}
	return fmt.Sprint(sum)

}
