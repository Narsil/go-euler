//Problem 26
//Find the value of d  1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.

package projecteuler

import (
	"fmt"
)

func Euler26() string {
	longest := 0
	for i := 1; i < 1000; i++ {
		a := 1
		b := i
		r := 1
		cycle := 0
		prev := 0
		presence := false
		quotients := make(map[int]int)
		for position := 0; r != 0; position++ {
			r = a % b
			prev, presence = quotients[r]
			if presence {
				cycle = position - prev + 1
				break
			}
			quotients[r] = position
			a = 10 * r
		}
		if cycle > longest {
			longest = cycle
		}
	}
	return fmt.Sprint(longest)
}
