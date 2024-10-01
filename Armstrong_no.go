package main
// fmt package provides the function to print anything
import "fmt"
func main() {
   fmt.Println("Number = 153")     // (1)3 + (5)3 + (3)3
   // declare the variables            1 +  125 + 27
                                     // 153
   var number, temp, rem int
   var result int = 0
   // initialize the variables
   number = 153
   temp = number
   // Use of For Loop
   for {
      rem = temp % 10
      result += rem * rem* rem
      temp /= 10
      if temp == 0 {
         break // Break Statement used to stop the loop
      }
   }
   // If satisfies Armstrong condition
   if result == number {
      fmt.Printf("%d is an Armstrong number.", number)
   } else {
      fmt.Printf("%d is not an Armstrong number.", number)
   }
   // print the result
}