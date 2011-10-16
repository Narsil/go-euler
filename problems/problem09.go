// Problem 9 - http://projecteuler.net/
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package projecteuler

import (
	"fmt"
)

func solve() (a, b, c int) {
	for a = 1; a < 1000; a++ {
		for b = a + 1; b < 1000; b++ {
			for c = b + 1; c < 1000; c++ {
				if a+b+c == 1000 && (a*a+b*b == c*c) {
					return a, b, c
				}
			}
		}
	}
	return 0, 0, 0
}

func Euler9() string {
	a, b, c := solve()
	return fmt.Sprint(a * b * c)
}
