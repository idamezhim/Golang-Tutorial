package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++{
	if i % 2 == 0 {
		fmt.Println(i , "divisible by 2") 
		}else if i % 3 == 0 {
		fmt.Println(i , "divisible by 3") 
	   } else if i % 4 == 0 {
		fmt.Println(i, "divisible by 4")
	   }else if i % 3 & 4 & 5 == 0 {
		   fmt.Println(i , "Divisible by 3,4,5")
	   }
	}
	
}