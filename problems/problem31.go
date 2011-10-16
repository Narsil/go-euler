//Problem 31
//How many different ways can £2 be made using any number of coins?

package projecteuler

import (
	"fmt"
)

func Euler31() string {
	count := 0
	var a, b, c, d, e, f, g int
	for a = 200; a >= 0; a -= 200 {
		for b = a; b >= 0; b -= 100 {
			for c = b; c >= 0; c -= 50 {
				for d = c; d >= 0; d -= 20 {
					for e = d; e >= 0; e -= 10 {
						for f = e; f >= 0; f -= 5 {
							for g = f; g >= 0; g -= 2 {
								count++
							}
						}
					}
				}
			}
		}
	}
	return fmt.Sprint(count)
}
