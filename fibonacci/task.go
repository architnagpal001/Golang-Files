package main

import "fmt"

func main(){
	fmt.Println(fact(3))
}
func fact(number int) int {
	if number <= 1 {
		return 1
	}
	var result int
	number == fact(number-1)
	return result
	
}