// Problem 1 - http://projecteuler.net/
//
// Add all the natural numbers below one thousand
// that are multiples of 3 or 5.

package projecteuler

import "fmt"

func Euler1() string {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	s := fmt.Sprint(sum)
	return s
}
