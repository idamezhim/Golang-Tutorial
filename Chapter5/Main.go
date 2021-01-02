package main

import "fmt"

func main() {
	x := make([]float64, 5)
	var total float64 = 0
for _, value := range x {
 total += value
}
fmt.Println(total / float64(len(x)))
   }