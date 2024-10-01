package main
import "fmt"
func main(){
	var x int =10
	var y int =20
	if(x >= 10){
	   if(y >= 10){
	     fmt.Printf("Inside Nested If Statement\n")
	   }
	}
	fmt.Printf("Value of x : %d\n",x)
	fmt.Printf("Value of y : %d\n",y)
}