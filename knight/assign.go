// // package main

// // var isKnightSleep bool
// // var isPrisonerSleep bool
// // var isArcherSleep bool
// // var ishasDog bool

// // func CanFastAttack() bool {
// // 	return isKnightSleep
// // }
// // func CanSpy() bool {
// // 	return !(isKnightSleep && isPrisonerSleep && isArcherSleep)
// // }
// // func CanSignalPrisoner() bool {
// // 	return (!isPrisonerSleep) && (isArcherSleep)
// // }
// // func CanFreePrisoner() bool{
// // 	return ((hasDog && isArcherSleep) || (!isPrisonerSleep))
// // }

// package main 
// import "fmt"

// func Hello(){
// 	var data int
//     fmt.Println("Enter your age=")
//     fmt.Scanf("%d", &data)
//     if data >= 18 {
//     fmt.Println("You are adult")
//     }else{
//     fmt.Println("You are a child")
//     }
// }
// func main(){
// 	Hello()
// 	//fmt.Println(data)
// }

package main

import "fmt"

type Pill int

const (
Placebo Pill = iota // iota =0,1,2,3,4,5
Aspirin //1
Ibuprofen //2
Paracetamol //3
Acetaminophen //4
)

/*
var(
name string
age int
)
*/
func main() {
fmt.Println(Paracetamol)
}
