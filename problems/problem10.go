// Problem 10 - http://projecteuler.net/
//
// Find the sum of all the primes below two million.

package projecteuler

import (
	"fmt"
	"primes"
)


func Euler10() string{
	sum := 0;
	for i:=0;i<2000000;i=<-primes.FastSieve() {
		sum += i
	}
	return fmt.Sprint(sum);
}
