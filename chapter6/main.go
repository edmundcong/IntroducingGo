package main

import ("fmt")

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(variadicFunction(xs...))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	//deferTest()
	//defer func() {
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("WAHEY!")
	v := 5
	mv := new(int)
	pointerTest(&v) // different ways to reference memory location
	pointerTest(mv)
	fmt.Println("new value for v is ", v)
	fmt.Println("value for mv is ", *mv)
}

func pointerTest(cv * int) {
	*cv = 10
}

func deferTest() (n int, err error) {
	defer fmt.Println("Defer")
	fmt.Println("Start")
	return fmt.Println("Return")
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func variadicFunction(args ...float64) float64{
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

// an example of a closure (anon func + inner func access to outer func variable)
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

