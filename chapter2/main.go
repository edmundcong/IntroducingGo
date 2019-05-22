package main

import "fmt"
import "math"

func main() {
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
	fmt.Println("5 % 3 =", 5 % 3)
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")
	fmt.Println("Hello \n world")
	fmt.Println("11111111 = ", math.Exp2(8)-1)
}