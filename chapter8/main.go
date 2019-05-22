package main

import (
	"fmt"
	"strings"
)


func main() {
	// some string functions to play around with
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	arr := []byte("test")
	fmt.Println(arr) // remember: in bytes!
	str := string(arr) // cast bytes slice to string
	fmt.Print((str))
}