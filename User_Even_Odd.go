package main
import "fmt"
func main(){
	var i int
	fmt.Printf("Enter Number:")
	fmt.Scanf("%d",&i)
	if(i%2==0){
	  fmt.Println("The Given Number is Even.")
	}else{
	  fmt.Println("The Given Number is Odd.")
	}
}