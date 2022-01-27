package main

import "fmt"

var i int
var j int

func main() {
	var row int
	var col int
	var m1 [10][10]int
	var m2 [10][10]int
	var sum [10][10]int

	fmt.Println("Enter the number of rows(should be >1 and <10): ")
	fmt.Scanln(&row)
	fmt.Println("Enter the number of rows(should be >1 and <10): ")
	fmt.Scanln(&col)
	fmt.Println("Enter the elements of 1st matrix: ")

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scanln(&m1[i][j])
		}
	}
	fmt.Println("Enter the elements of 2nd matrix: ")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scanln(&m2[i][j])
		}
	}

	fmt.Println("Output: ")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = m1[i][j] + m2[i][j]
			fmt.Println(sum[i][j])
		}
	}

}
