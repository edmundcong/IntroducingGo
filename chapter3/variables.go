package main

import "fmt"

var outerScope string = "variable which is on the outer scope. Also camelCasing"

func main() {
	var x string = "Hello, world"
	fmt.Println(x)
	x += " and then some"
	fmt.Println(x)
	y := "this is going to be a string"
	fmt.Println(y)
	latterFunction()
	fmt.Println("Enter celcius: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := 32.0 + (input / (5.0 / 9.0))
	fmt.Println("in fahrenheit that is: ", output)
}

func latterFunction() {
	const c string = "some const"
	fmt.Println(c)
}

