//Problem 21
// Evaluate the sum of all amicable pairs under 10000

package projecteuler

import (
	"fmt"
	"primes"
)

func Euler21() string {
	sum := uint64(0)
	for i := uint64(0); i < uint64(10000); i++ {
		pair := primes.SumDivisors(i)
		if i == primes.SumDivisors(pair) && i != pair {
			sum += i
		}
	}
	return fmt.Sprint(sum)
}
