package main

import "fmt"

var s1 string
var c string

func main() {
	fmt.Println("Enter a string:")
	fmt.Scanln(&s1)
	checkIso()
}
func checkIso() bool {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if s1[i] != s1[j] {
				c = "false"
			}
		}
	}
	fmt.Println(c)
	return true
}
