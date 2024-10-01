package main
 
import "fmt"
 
// declaring a function
// having return type
// of int, int
func maxmin(a int, b int) (int, int) {
 
    if a > b {
     
        // separate multiple return
        // values using comma
        return a, b
    } else {
 
        return b, a
    }
    // this function returns
    // maximum , minimum values
}
 
func main() {
 
    // declaring two values a and b
    var a = 50
    var b = 70
 
    // calling the function
    // with multiple assignments
    var max, min = maxmin(a, b)
 
    // Printing the values
    fmt.Println("Max = ", max, "\nMin = ", min)
}

