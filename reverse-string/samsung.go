package main

import "fmt"

var str string

func reverse() (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println("Enter the string you want to reverse :")
	fmt.Scanf("%s", &str)

	// strRev := reverse(st)
	var strRev string
	fmt.Println(str)
	fmt.Println(strRev)

}
