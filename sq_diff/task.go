// package main
// import "fmt"

// func main() {
// 	var n int
// 	defer squareofsum(n)
//  //Enter your code here. Read input from STDIN. Print output to STDOUT
//     fmt.Println("Enter Value of n:")
// 	fmt.Scanf("%d", &n)
// }

// func squareofsum(n int) {
//     var s1 int
//     var sq int
//     s1 = (n*(n+1))/2
//     sq = s1 * s1
//     fmt.Println(sq)
// }

// package main
// import "fmt"

// func main() {
//  var i int
//     fmt.Scanf("%d", &i)
//     fmt.Println(Difference(i))
// }

// func SquareOfSum(n int) int {
//     var res int
//     //sum all the numbers
//     for i := 1; i <= n; i++ {
//         res = res+i
//     }
//     //square them
//     res = res*res
//     return res
// }

// func SumOfSquares(n int) int {
//     var res int
//     //square and sum them together
//     for i := 1; i <= n;i++ {
//         temp := i*i
//         res = res+temp
//     }
//     return res
// }

// func Difference(n int) int {
//     var squaresum, sumsquare int
//     squaresum = SquareOfSum(n)
//     sumsquare = SumOfSquares(n)
//     res := squaresum - sumsquare
//     return res   
// }

package main
import "fmt"

func main() {
 var a int
    fmt.Println("Enter the Nummber:")
    fmt.Scanf("%d", &a)
    fmt.Println(final(a))
}

func Sq_sum(n int) int {
    var result int
    for i := 1; i <= n; i++ {
        result = result+i
    }
    result = result * result
    return result
}

func Sum_sq(n int) int {
    var result int
    for j := 1; j <= n;j++ {
        temp := j*j
        result = result + temp
    }
    return result
}

func final(n int) int {
    var k, l int
    k = Sq_sum(n)
    l = Sum_sq(n)
    result := k - l
    return result
    
}