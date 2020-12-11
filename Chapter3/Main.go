package main
import "fmt"

func main() {
	var x string
 	x = "first"
 	fmt.Println(x)
 	x = "second"
	fmt.Println(x)
	x = "hello"
	var y string = "world"
	fmt.Println(x == y)
	y = "world"
	fmt.Println(x == y)
	y = "hello"
	fmt.Println(x == y)
	t := "football"
	r := "No football"
	fmt.Println(t == r)
	x = "football"
	fmt.Println(x == t)
	g := 55
	fmt.Println(g)
	var(
		a = 2
		b = 3
		c = 4
	)
	fmt.Println(a + b + c)
	const i int = 345
	fmt.Println(i + c)
	fmt.Println(i == b)
	//This is a Program to convert Celsuis to Fahre heit
	fmt.Print("Enter the Fahrenheit Degree: ")
	var f float64
	fmt.Scanf("%f", &f)
	output :=  (f - 32) * 5/9
	fmt.Println(output)
} 