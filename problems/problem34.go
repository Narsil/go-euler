//Problem 34
//Find the sum of all numbers which are equal to the sum of the factorial of their digits.

package projecteuler

import (
	"fmt"
	"strconv"
	"big"
)

func Euler34() string {
	sum := 0
	for a := 10; a < 100000; a++ { // lowest power of ten with the good answer,guesswork
		digits := strconv.Itoa(a)
		sumFact := big.NewInt(0)
		fact := big.NewInt(0)
		for _, char := range digits {
			digit, _ := strconv.Atoi(string(char))
			sumFact = sumFact.Add(sumFact, fact.MulRange(1, int64(digit)))
		}
		if sumFact.Int64() == int64(a) {
			sum += a
		}
	}
	return fmt.Sprint(sum)
}
