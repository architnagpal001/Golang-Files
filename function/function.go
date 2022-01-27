package main
import "fmt"

func sum(x int,y int){
	var data1 int
	var data2 int
	fmt.Println("Enter a value =")
	fmt.Println("Enter 2nd value =")
	fmt.Scanf("%d", &data1)
	fmt.Scanf("%d", &data2)
	var sum = x+y
	return sum
}

// func main(){
// 	var x int
// 	var y int
// 	fmt.Println("enter 1st value")
// 	fmt.Scanf("%d", x)
// 	fmt.Println("enter 2nd value")
// 	fmt.Scanf("%d", y)
// 	sum = x+y
// 	fmt.Println(sum(x+y))
// }