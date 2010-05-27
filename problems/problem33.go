//Problem 33
//Discover all the fractions with an unorthodox cancelling method.

package projecteuler

import (
	"fmt"
	"strconv"
	"exp/bignum"
)

func hasProperty( a,b int ) bool{
	stra := strconv.Itoa(a)
	strb := strconv.Itoa(b)
	num := uint8(0)
	den := uint8(0)
	if stra[0]==strb[0] {
		num = stra[1]
		den = strb[1] 
	}else if stra[0] == strb[1]{
		num = stra[1]
        den = strb[0] 
	}else if stra[1] == strb[1] && stra[1]!=uint8(48){
		num = stra[0]
        den = strb[0] 
	}else if stra[1] == strb[0]{
		num = stra[0]
        den = strb[1] 
	}else{
		return false
	}
	newa,_ := strconv.Atof(string(num))
    newb,_ := strconv.Atof(string(den))
    if float(a)/float(b) == newa/newb{
		return true
	}
	return false
}

func gcd (u,v int) int{
  if u < 0 {
        return gcd(-u, v)
    }
    if u == 0 {
        return v
    }
    return gcd(v%u, u)
}

func Euler33()string{
	prod:=bignum.Rat(1,1)
	for a:=10;a<100;a++{
		for b:=a+1;b<100;b++{
			if hasProperty(a,b){
				prod= prod.Mul(bignum.Rat(int64(a),int64(b)))
			}
		}
	}
	_,den := prod.Value()
	return fmt.Sprint(den.String())
}
