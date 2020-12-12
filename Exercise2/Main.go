package main
import "fmt"

func main() {
	//This is a Program to convert Celsuis to Fahreheit
	fmt.Print("Enter the Celsuis Degree: ")
	var c float64
	fmt.Scanf("%f", &c)
	output :=  (c * 9/5) + 32
	fmt.Println(output)
}
