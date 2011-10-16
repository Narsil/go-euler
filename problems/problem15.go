//Problem 15
//Starting in the top left corner in a 20 by 20 grid, how many routes are there to the bottom right corner?

package projecteuler

import (
	"fmt"
	"big"
)

func Euler15() string {
	//We need to choose Right or Down, 40 times, but we need
	//to choose 20 Right and 20 Down (no backtracking).
	//This number is just 20C40, expressed with bigs
	//to avoid overflow
	fact := big.NewInt(0)
	fact.MulRange(1, 20)
	fact.Mul(fact, fact)

	bigfact := big.NewInt(0)
	bigfact.MulRange(1, 40)
	return fmt.Sprint(bigfact.Div(bigfact, fact))
}
