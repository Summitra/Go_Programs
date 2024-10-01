package main
import "fmt"
func main(){
	var int a
	var int b
	fmt.Printf("Enter First String:")
	fmt.Scanf("%d",&a)
	fmt.Printf("\nEnter the Second String:")
	fmt.Scanf("%d",&b)
	fmt.Println("String is:", string.join(a+b))
}