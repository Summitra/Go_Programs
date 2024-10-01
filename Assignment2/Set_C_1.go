
package main

import "fmt"

func Swap(num1 int, num2 int) {
	var temp int = 0

	temp = num1
	num1 = num2
	num2 = temp
}
func main() {
	var num1 int = 10
	var num2 int = 20

	fmt.Println("Numbers before swapping: ", num1, num2)
	Swap(num1, num2)
	fmt.Println("Numbers after swapping: ", num1, num2)
}