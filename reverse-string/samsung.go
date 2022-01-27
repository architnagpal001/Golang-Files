package main

import "fmt"

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var st string
	fmt.Println("Enter the string you want to reverse :")
	fmt.Scanf("%s", &st)

	strRev := reverse(st)
	fmt.Println(st)
	fmt.Println(strRev)

}
