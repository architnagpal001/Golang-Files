// package main 
// import "fmt"

// func main(){
// 	var data int
// 	fmt.Println("Enter a Value")
// 	fmt.Scanf("%d", &data)

// 	if data < 18 {
// 		fmt.Println("You are not adult")
// 	} else{
// 		fmt.Println("You are adult")
// 	}
// }


package main 
import "fmt"

func main(){
	var i int
	var data int
	for i=0 ; i<3 ; i++{
		fmt.Println("Enter a Value")
		fmt.Scanf("%d", &data)

		if data < 18 {
			fmt.Println("You are not adult")
			break
		} else{
			fmt.Println("You are adult")
		}
	}
}