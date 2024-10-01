package main
import "fmt"

func main() {
	var num = 1 

	fmt.Printf("Enter Number: ")
	fmt.Scanf("%d", &num)

	for i:= 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}