// Problem 36 - http://projecteuler.net/
//
// Find the sum of all numbers, less than one million, 
// which are palindromic in base 10 and base 2.

package projecteuler

import (
	"fmt"
	"strconv"
)

func Euler36() string {
	sum := 0
	for i := 1; i < 1000000; i++ {
		if palidrome(strconv.Itoa(i)) &&
			palidrome(strconv.Itob(i, 2)) {
			sum += i
		}
	}
	return fmt.Sprint(sum)
}
