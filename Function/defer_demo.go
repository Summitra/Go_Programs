package main
import "fmt"
func mult(a int, b int)int{
	return a*b
}
func show(){
	fmt.Println("Welcome")
}
func main(){
	show()
	defer mult(10,30)
	mult (20,60)
}