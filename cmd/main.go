package main

import (
	"projecteuler"
	"fmt"
	"profiling"
)

func main(){
	for i:=20;i<37;i++ {
		if i==10 || i==23 || i==25 || i==27{
			//problems takes forever, needs optimization
			continue
		}
		time, result := profiling.TimeArg(projecteuler.Euler,i)
		fmt.Printf("Problem %v : %v\t%v s\n",i,result, time/1000000000.0)
	}
}
