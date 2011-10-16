// Problem 41
//What is the largest n-digit pandigital prime that exists?

package projecteuler

import (
	"fmt"
	"big"
	"arithmetics"
	"strconv"
)

func Euler41() string {
	maxPandigital := "123456789"
	max := 0
	for i := len(maxPandigital); i > 3; i-- {
		pandigital := []byte(maxPandigital[0:i])
		for arithmetics.Permute(pandigital) {
			number, _ := strconv.Atoi(string(pandigital))
			if big.ProbablyPrime(big.NewInt(int64(number)), 1) {
				if number > max {
					max = number
				}
			}
		}
		if max != 0 {
			return fmt.Sprint(max)
		}
	}
	return "Error not prime found"
}
