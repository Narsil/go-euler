// Problem 25 - http://projecteuler.net/
//
// What is the first term in the Fibonacci sequence to contain 1000 digits?

package projecteuler

import (
	"fmt"
	"exp/bignum"
)

func Euler25() string{
	a := bignum.Int(1)
	b := bignum.Int(1)
	i := 2
	for {
		a, b = b, a.Add(b)
		i++
		if len(b.String()) == 1000 {
			break
		}

	}
	return fmt.Sprint(i)
}
