package main
import "fmt"
func main(){
	var Str string
	fmt.Printf("Enter the String:")
	fmt.Scanf("%s",&Str)
	fmt.Println("The Entered String is:",len(Str))// Length is calculated from index no 1
}