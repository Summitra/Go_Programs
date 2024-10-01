package main
import "fmt"
func main(){
	fact()
}
func fact(){
	var f int = 1
	var n int
	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &n  )
	if(n>0){
	f= f*n
	n--
	fmt.Println("factorial of a number is :", fact())
	}
}