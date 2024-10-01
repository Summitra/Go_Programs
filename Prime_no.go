package main

import "fmt"

func main() {
	var num int = 0
	var flag int = 0

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	for count := 2; count < num/2; count++ {
		if num%count == 0 {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Print("Given number is not PRIME number")
	} else {
		fmt.Print("Given number is PRIME number")
	}
}