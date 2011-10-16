package bigextension

import (
	"big"
)

func Pow(x *big.Int, y int64) *big.Int {
	var z *big.Int
	if y == 1 {
		return x
	}
	var q, r int64
	q = y / 2
	r = y % 2

	z = Pow(x, q)
	z.Mul(z, z)

	if r == 0 {
		return z
	}
	return z.Mul(z, x)

}
