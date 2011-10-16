//Problem 23
//Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

package projecteuler

import (
	"fmt"
	"primes"
)

func Euler23() string {
	//Slow, bruted forced any more elegant solution welcome.
	sum := uint64(0)
	count := 0
	abundants := make([]int, 8000)
	sumAbundants := make(map[int]int)
	for i := 1; i < 28124; i++ {
		if primes.SumDivisors(uint64(i)) > uint64(i) {
			abundants[count] = i
			count++
		}
	}
	count--
	curSum := 0
	for i := 0; i < count; i++ {
		for j := i; j < count; j++ {
			curSum = abundants[i] + abundants[j]
			_, present := sumAbundants[curSum]
			if curSum <= 28123 && !present {
				sumAbundants[curSum] = curSum
			}
		}
	}
	for i := 1; i < 28123; i++ {
		_, present := sumAbundants[i]
		if !present {
			sum += uint64(i)
		}
	}
	return fmt.Sprint(sum)
}
