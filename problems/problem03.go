// Problem 3 - http://projecteuler.net/
//
// Find the largest prime factor of a composite number

package projecteuler

import (
	"primes"
	"fmt"
)

func Euler3() string {
	const composite = 600851475143
	factors := primes.Factors(composite)
	largest := factors[len(factors)-1]
	s := fmt.Sprint(largest)
	return s
}
