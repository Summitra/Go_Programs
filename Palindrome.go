// Golang program to check the given number is palindrome or not
// using the goto statement
package main
import "fmt"

func main() {
	var num int = 0
	var rev int = 0
	var rem int = 0
	var temp int = 0

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	temp = num
MyLbl:

	rem = num % 10
	rev = rev*10 + rem
	num = num / 10

	if num > 0 {
		goto MyLbl
	}

	if temp == rev {
		fmt.Printf("Number is palindrome")
	} else {
		fmt.Printf("Number is not palindrome")
	}
}