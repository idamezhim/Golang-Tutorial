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
func f() {
	//This is a Program to convert Fahreheit to Celsuis
	fmt.Print("Enter the Celsuis Degree: ")
	var c float64
	fmt.Scanf("%f", &c)
	output :=  (-5 * 9/5) + 32
	fmt.Println(output)
}