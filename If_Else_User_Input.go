package main
import "fmt"
func main(){
	var a int
	fmt.Print("Enter Number:")
	fmt.Scanln(&a)
	fmt.Print(a)
	if(a % 2 ==0){
	   fmt.Printf(" is Even Number\n")
	}else{
	   fmt.Printf(" is Odd Number\n")
	}
}