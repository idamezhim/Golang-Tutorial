package main

import "fmt"

func main(){
  for i := 1; i <= 100; i++{
		if  i % 3 == 0 {
			fmt.Println( i, "divisible by 3")
		 } else{
			 fmt.Println(i, "Not divisible by 3")}
							}
			}