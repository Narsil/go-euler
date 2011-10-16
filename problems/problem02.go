// Problem 2 - http://projecteuler.net/
//
// Find the sum of all the even-valued terms in Fib
// the sequence which do not exceed four million.

package projecteuler

import "fmt"

func Euler2() string {
	sum := 0
	for a, b := 1, 2; b < 4000000; {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, b+a
	}
	s := fmt.Sprint(sum)
	return s
}
