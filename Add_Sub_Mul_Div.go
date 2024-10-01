package main
import "fmt"
func main(){
	var a,b,c,ch int
	fmt.Println("Enter the First No:")
	fmt.Scanln(&a)
	fmt.Println("Enter the Second No:")
	fmt.Scanln(&b)
	fmt.Println("Enter the Choice:")
	fmt.Scanln(&ch)
	switch(ch){
	case 1: c= a+b
	        fmt.Println("Addition of Numbers is: ",c)
	break
	case 2: c=a-b
	        fmt.Println("Subtraction of Numbers is:",c)
	break
	case 3: c= a*b
	        fmt.Println("Multiplication of Numbers is:",c)
	break
	case 4: c= a/b
	        fmt.Println("Division of Numbers is:",c)
	break
	default : fmt.Println("Invalid Input !!")
	}
}