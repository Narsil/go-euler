//Problem 14
//Find the longest sequence using a starting number under one million.

package projecteuler

import (
	"fmt"
)

func flight(n uint64) uint64 {
	cur := n
	length := uint64(0)
	for cur != 1 {
		if cur%2 == 0 {
			cur /= 2
		} else {
			cur = 3*cur + 1
		}
		length++
	}
	return length
}

func Euler14() string {
	max := uint64(0)
	hit := uint64(0)
	for i := uint64(1); i < 1000000; i++ {
		tmp := flight(i)
		if tmp > max {
			max = tmp
			hit = i
		}
	}
	return fmt.Sprint(hit)
}
