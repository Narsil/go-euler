//Problem 15
//Starting in the top left corner in a 20 by 20 grid, how many routes are there to the bottom right corner?

package projecteuler

import (
	"fmt"
	"exp/bignum"
)

func Euler15() string{
	//We need to choose Right or Down, 40 times, but we need
	//to choose 20 Right and 20 Down (no backtracking).
	//This number is just 20C40, expressed with bignums
	//to avoid overflow
	return fmt.Sprint(bignum.Fact(40).Div(bignum.Fact(20).Pow(2)))
}
