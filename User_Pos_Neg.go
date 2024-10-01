package main
import "fmt"
func main(){
	var i int
	fmt.Printf("Enter the Number:")
	fmt.Scanf("%d",&i)
	if(i>0){
	   fmt.Println("Number is Positive")
	}else if(i<0){
	  fmt.Println("Number is Negative")
	}else{
	  fmt.Println("Number is zero")
	}
}