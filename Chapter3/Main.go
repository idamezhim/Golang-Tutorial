package main
import "fmt"

func main() {
	fmt.Println(len("Hello, World"))
    fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")
	fmt.Println(true)
 	fmt.Println(true && false)
 	fmt.Println(true)
 	fmt.Println(true || false)
	fmt.Println(!true)
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

	} 