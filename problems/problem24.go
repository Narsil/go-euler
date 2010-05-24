//Problem 24
//What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

package projecteuler

import (
	"fmt"
)

func swap(s []byte, a,b int) {
	s[a],s[b] = s[b],s[a]
}

func permute(str []byte) int{
	length := len(str)
	key := length-1
	newkey := key
	for ;key > 0 && str[key] <= str[key-1];{
		key--
	}
	key--
	if key < 0{
		return 0
	}
	newkey = length-1
	for ;newkey > key && str[newkey] <= str[key];{
		newkey--
	}
	swap(str,key,newkey)
	length--
	key++
	for ;length>key;{
		swap(str,length,key)
		key++
		length--
	}
	return 1
}

func Euler24()string{
	str := "0123456789"
	bytes := []byte(str)
	for i:=1;i<1000000;i++{
		permute(bytes)
	}
	return fmt.Sprint(string(bytes))
}
