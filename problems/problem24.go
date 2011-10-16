//Problem 24
//What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

package projecteuler

import (
	"fmt"
	"arithmetics"
)

func Euler24() string {
	str := "0123456789"
	bytes := []byte(str)
	for i := 1; i < 1000000; i++ {
		arithmetics.Permute(bytes)
	}
	return fmt.Sprint(string(bytes))
}
