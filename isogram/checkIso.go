package main

import "fmt"

var s1 string
func main(){
	fmt.Println("Enter a string:")
	fmt.Scanln(&s1)
}
func checkIso() bool{
	for i:=0 ; i < len(s1) ; i++ {
		for j := 0 ; j < len(s1) ; j++ {
			if (s1[i] == s1[j]) {
				return false
			}
		}
		fmt.Println("It is isogram")
	}
	return true
}