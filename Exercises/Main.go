package main
import "fmt"

func main() {
	//This is a Program to convert Celsuis to Fahreheit
	fmt.Print("Enter the Fahrenheit Degree: ")
	var f float64
	fmt.Scanf("%f", &f)
	output :=  (f - 32) * 5/9
	fmt.Println(output)
}