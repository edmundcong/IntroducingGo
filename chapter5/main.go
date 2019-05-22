package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var results [10]int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano()) // dynamic seed so we get different rand each run
		results[i] = rand.Intn(100)
		fmt.Println("results[", i, "] = ", results[i])
	}
	sum := 0.0
	for _, value := range results{
		sum += float64(value)
	}
	res := sum/float64(len(results))
	fmt.Println("result = ", res)
	// anotherWayToCreateArrays := [5]float64{1,2,3,4,5}
	mapsExample()
	fmt.Println(sliceExample())
}

func mapsExample() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	_, ok := elements["li"]
	fmt.Println(ok)

}

func sliceExample() []float64{
	arr := [5]float64{1,2,3,4,5}
	slce := arr[0:4]
	return slce
}