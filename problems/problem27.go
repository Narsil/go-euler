//Problem 27
//Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.

package projecteuler

import (
	"fmt"
	"big"
)

func Euler27() string {
	max := 0
	prod := 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			n := 0
			for ; big.ProbablyPrime(big.NewInt(int64(n*n+a*n+b)), 1); n++ {
			}
			if n > max {
				max = n
				prod = a * b
			}
		}
	}
	return fmt.Sprint(prod)
}
