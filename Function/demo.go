package main
import "fmt"
func main(){
	add()
}
func add(){
	var a,b,c int
	fmt.Println("Enter any two numbers: ")
	fmt.Scanf("%d %d", &a, &b)
	c = a+b
	fmt.Printf("Addition of two numbers is : %d",c)
}