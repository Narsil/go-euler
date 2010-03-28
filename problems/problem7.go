// Problem 7 - http://projecteuler.net/
//
// What is the 10001st prime number?

package projecteuler

import (
	"fmt"
	"primes"
)

func Euler7() string{
	sieve:=primes.FixedSieve(10001)
	i:=uint64(0)
	num:=uint64(0)
	for {
		num=i
		i=<-sieve
	}
	return fmt.Sprint(num)
}
