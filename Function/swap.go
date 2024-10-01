package main
import "fmt"
func main(){
	var a int 
	var b int 
	fmt.Printf("Enter two numbers:")
	fmt.Scanf("%d %d ",&a, &b)
	fmt.Printf("Before Swapping, Value of a =%d , b = %d\n",a,b)
	swap(a,b)
	fmt.Printf("After Swapping , Value of a = %d , b= %d\n", a,b)
}
func swap(x,y int)int{
	var temp int
	temp = x
	x=y
	y=temp
	return temp
}