// Problem 7 - http://projecteuler.net/
//
// What is the 10001st prime number?

package projecteuler

import (
	"fmt"
	"primes"
)

func Euler7() string {
	limit := 10001
	sieve := primes.FastSieve()
	num := 0
	for i := 0; i < limit; i++ {
		num = <-sieve
	}
	return fmt.Sprint(num)
}
