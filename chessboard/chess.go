package main

import "fmt"

var chessBoard [8][8]string

func checkOccupiedfiles(file int) int {
	count := 0
	for index1, _ := range chessBoard {
		if chessBoard[index1][file] == "#" {
			count++
		}
	}
	return count
}

func countOccupied() int{
	count:=0
	for index, _ := range chessBoard{
		for _, value := range chessBoard[index] {
			if value == "#" {
				count++
			}
		}
	}
	return count
}

func countAll() int{
	count := 0 
	for index , _ : range chessBoard {
		for _ , value : range chessBoard[index] {
			count++
		}
	} 
}


func main() {
	chessBoard = [8][8]string{
		{"#", "_", "#", "#", "_", "_", "_", "_"},
		{"#", "_", "#", "#", "_", "_", "_", "_"},
		{"#", "#", "_", "_", "#", "#", "#", "#"},
		{"#", "_", "#", "#", "_", "_", "_", "#"},
		{"#", "_", "#", "#", "_", "_", "_", "#"},
		{"#", "_", "#", "#", "_", "_", "_", "#"},
		{"#", "_", "#", "#", "_", "_", "_", "#"},
		{"#", "_", "#", "#", "_", "_", "_", "#"},
	}
	fmt.Println(checkOccupiedfiles(2))
	fmt.Println(countOccupied())
	fmt.Println(countAll())
}