package main

import (
	"errors"
	"fmt"
)

const pr = 221

func successrate(i int) (float64) {
	if i == 0 {
		return 0
	}
	if i == 1 || i == 2 || i == 3 || i == 4 {
		return 1.0
	}
	if i == 5 || i == 6 || i == 7 || i == 8 {
		return 0.9
	}
	if i == 9 || i == 10 {
		return 0.77
	}
	return 0
	//retruning the success rate 
}
func prph(i int) float64 {
	return successrate(i) * float64(i) * pr
}
func prpm(i int) int {
	return int(prph(i)) / 60
}
func alp(i, a int) int {
	if int(prph(i)) <= a {
		return int(prph(i))
	} else {
		return a
	}
}
func main() {
	var i int
	var a int
	fmt.Print("enter the speed =")
	fmt.Scan(&i)
	fmt.Print("enter artificial limit=")
	fmt.Scan(&a)
	fmt.Println("the success rate is --------------->", s)
	fmt.Println("the production rate per hour is --->", prph(i))
	fmt.Println("the production rate per minute is ->", prpm(i))
	fmt.Println("the artificially limited rate is -->", alp(i, a))
	}
}
