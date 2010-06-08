//Problem 44
//Find the smallest pair of pentagonal numbers whose sum and difference is pentagonal.

package projecteuler

import (
	"fmt"
	"math"
)

func isPentagonal(n int)bool{
  sq := math.Sqrt(float64(1+24*n))
  return sq*sq == float64(1+24*n) && int(1+sq) % 6 == 0
}

func Euler44() string{
	pentagonals:=make(map[int]int)
	for i:=1;;i++{
		pent1 := i*(3*i-1)/2
		for pent2 := range pentagonals{
			diff := pent1-pent2
			_,presence := pentagonals[diff]
			if presence{
				sum := pent1+pent2
				if isPentagonal(sum){
					fmt.Println(diff,pent1,pent2,sum)
					return fmt.Sprint(diff)
				}
			}
		}
		pentagonals[pent1]=1
	}
	return "Error occured"
}
