// Problem 12 - http://projecteuler.net/
//

// What is the value of the first triangle number
// to have over five hundred divisors?

package projecteuler

import (
	"fmt"
	"primes"
)


func Euler12() string {
	triangle := uint64(1)
	for i:=uint64(2); ;i++{
		c := primes.NumDivisors(triangle)
		if c>500{
			break
		}
		triangle+=i
	}
	return fmt.Sprint(triangle);
}
