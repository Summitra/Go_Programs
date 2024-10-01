package main
import (
	"fmt"
)
func swap(a *int, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}
func main() {
	var a int
	var b int
	fmt.Printf("Enter First number : ")
	fmt.Scanln(&a)
	fmt.Printf("Enter Second number : ")
	fmt.Scanln(&b)

	fmt.Printf("Before swap: First = %d, Second = %d \n", a, b)
	// Function Call
	
	swap(&a, &b)
	fmt.Printf("After swap: First = %d, Second = %d \n", a, b)

}