//Problem 20
// Find the sum of the digits of 100!

package projecteuler

import (
	"fmt"
	"exp/bignum"
	"strconv"
)

func Euler20() string{
	str := bignum.Fact(100).String()
	num:=0
	digit:=0
	for _,char :=range str{
		digit,_ = strconv.Atoi(string(char))
		num += digit
	}
	return fmt.Sprint(num)
}
