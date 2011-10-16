//Problem 38
//What is the largest 1 to 9 pandigital that can be formed by multiplying a fixed number by 1, 2, 3, ... ?

package projecteuler

import (
	"fmt"
	"strconv"
)

func Euler38() string {
	max := 0
	for i := 0; i < 10000; i++ {
		str := ""
		for j := 1; j < 10; j++ {
			str += strconv.Itoa(j * i)
			if len(str) > 10 {
				break
			}
			if pandigital(str) {
				num, _ := strconv.Atoi(str)
				if num > max {
					max = num
				}
			}
		}
	}
	return fmt.Sprint(max)
}
