package main

import (
	"projecteuler"
	"fmt"
	"profiling"
)

func main(){
	for i:=1;i<8;i++ {
		time, result := profiling.TimeArg(projecteuler.Euler,i)
		fmt.Printf("Problem %v : %v\t%v s\n",i,result, time/1000000000.0)
	}
}
