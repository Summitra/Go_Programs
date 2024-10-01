package main

import "fmt"

func addNumber(number1, number2 int) int {
   return number1 + number2
}
func main() {

   var number1, number2, number3 int
   
   // initializing the variables
   number1 = 18
   number2 = 9
   
   number3 = addNumber(number1, number2)
   fmt.Println("The Addition of", number1 ,"and" , number2 ,"is :", number3)
}