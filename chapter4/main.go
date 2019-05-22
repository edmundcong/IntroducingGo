package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
	for j := 1; j <= 10; j++{
		switchExample(j)
	}
}

func switchExample(i int) {
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	case 6: fmt.Println("Six")
	case 7: fmt.Println("Seven")
	case 8: fmt.Println("Eight")
	case 9: fmt.Println("Nine")
	default: fmt.Println("Unknown")
	}
}