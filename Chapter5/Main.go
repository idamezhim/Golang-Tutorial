package main

import "fmt"

func main(){
	var x [19]int
	x[4] = 100
	x[12] = 23
	x[3] = 24
	x[2] = 34
	fmt.Println(x[4], x[12], x[3], x[2], x)
}