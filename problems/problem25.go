// Problem 25 - http://projecteuler.net/
//
// What is the first term in the Fibonacci sequence to contain 1000 digits?

package projecteuler

import (
	"fmt"
	"big"
)

func Euler25() string{
	a := big.NewInt(1)
	b := big.NewInt(1)
    tmp := big.NewInt(0)
	i := 2
	for {
		tmp.Add(a, b)
        b = a
        a = tmp
		i++
		if len(b.String()) == 1000 {
			break
		}

	}
	return fmt.Sprint(i)
}
