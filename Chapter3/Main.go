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
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
	o := 5;
	o += 1
} 