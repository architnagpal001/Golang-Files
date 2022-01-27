// package main
// import "fmt"

// func main(){
// 	var a string
// 	a = "GAGCCTACTAACGGGAT"
// 	var b string
// 	b = "CATCGTAATGACGGCCT"
// 	var c int
// 	for i:= range a {
// 		if a[i] != b[i]{
// 			c = c+1
// 		}
// 	}
// 	fmt.Println("The Hamming distance = ",c)
// }

package main
import "fmt"

func main(){
	var a string
	fmt.Println("Enter the Value of 1st Strand = ")
	fmt.Scanln(&a)
	var b string
	fmt.Println("Enter the Value of 2nd Strand = ")
	fmt.Scanln(&b)
	var c int
	for i:= range a {
		if a[i] != b[i]{
			c++
		}
	}
	fmt.Println("The Hamming distance = ",c)
}
