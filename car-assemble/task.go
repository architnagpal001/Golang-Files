// package main
// import "fmt"

// func successrate() int{
// 	var speed int
// 	if speed == 0{
// 		return 0
// 	}
// 	if speed >= 1 && speed<=4 {
// 		return 100
// 	}
// 	if speed >= 5 && speed <=8 {
// 		return 90
// 	}
// 	if speed >= 9 && speed <=10 {
// 		return 77
// 	} else{
		
// 	}
// }

// func rateperhour(a){
// 	successrate()
// 	var rph int 
// 	rph = 221 * speed * a
// 	fmt.Println(rph)
// 	return r*successrate(speed)
// }

// // func rateperminute() int{
// // 	rateperhour(speed)/60
// // }

// // func main(){
// // 	successrate()
// // 	rateperhour()
// // 	rateperminute()
// // 	fmt.Println("Enter the speed =")
// // 	fmt.Scanf("%d", &speed)
// // 	a := successrate(speed)
// // 	b := rateperhour(a)
// // 	c := rateperminute(b)
// // }

package main 

import "fmt"

func successrate(speed int) int{
	if speed == 0{
		return 0
	}
	if speed >= 1 && speed<=4 {
		return 100
	}
	if speed >= 5 && speed <=8 {
		return 90
	}
	if speed >= 9 && speed <=10 {
		return 77
	} else{
		fmt.Println("The Speed you entered is invalid")
		return 0
	}
}

func main() {
	var speed int
	var x int
	fmt.Println("Enter the speed =")
	fmt.Scanf("%d", &speed)
	x = successrate(speed)
	fmt.Println(x)
}