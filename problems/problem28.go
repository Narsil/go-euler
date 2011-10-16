//Problem 28
//What is the sum of both diagonals in a 1001 by 1001 spiral?

package projecteuler

import (
	"fmt"
)

func Euler28() string {
	sum := 1
	for i := 3; i <= 1001; i += 2 {
		sum += 4*i*i - 6*i + 6
	}
	return fmt.Sprint(sum)
}
