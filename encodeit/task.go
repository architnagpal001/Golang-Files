package main
import (
    "fmt"
    "strconv"
       )

var str string
var let int
var s string
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 	fmt.Println("Enter the string you want to encode :")
    fmt.Scanf("%s", &s)
    str := EncodeIt(s)
    fmt.Println(str)
    
}
func EncodeIt(s1 string) string {
    for i := 0; i<len(s1); i++ {
        let = 1
        for ;i<len(s1)-1 && s1[i] == s1[i+1];{
            let = let+1
            i=i+1
        }
        if let == 1 {
            str = str + string(s1[i])
            
        } else {
            str  =  str + strconv.Itoa(let) + string(s1[i])
            //lets do it
        }
        
    }
    return str
}
