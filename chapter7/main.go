package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func main() {
	l := Circle{0.0,0.0, 4.0}
	fmt.Println(l)
	fmt.Println(l.x, l.y, l.r)
	c := &Circle{0.0, 0.0, 5.0}
	fmt.Println(c)
	/*
	* pointers to struct's fields are automatically dereferenced!!!!!
	* c.x === (*c).x
	*/
	fmt.Println(c.x, c.y, c.r)
	fmt.Println((*c).x, (*c).y, (*c).r)

	cc := &Circle{ 1.5, 1.5, 3}
	fmt.Println(cc.area()) // no need to pass address in (i.e. &cc.area()) since go automatically infers this with receivers
}