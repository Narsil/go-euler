// Problem 10 - http://projecteuler.net/
//
// Find the sum of all the primes below two million.

package projecteuler

import (
	"fmt"
	"primes"
)


func Euler10() string{
	sum := uint64(0);
	for i:=uint64(0);i<2000000;i=<-primes.Sieve() {
		sum += i
	}
	return fmt.Sprint(sum);
}
