package main
import "fmt"
func main(){
	var a,b int
	fmt.Println("Enter any two numbers:")
	fmt.Scanf("%d %d", &a, &b)
    add(a,b)
}
func add(x,y int )int{
	var c int
	c = x+y
	fmt.Println("Addition of two numbers is :",c)
	return c
}