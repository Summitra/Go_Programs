package  main
import "fmt"
func main(){
	fmt.Print("Enter Number :")
	var a int
	fmt.Scanln(&a)
	switch(a){
	case 10:
	   fmt.Println("Value is 10")
	case 20:
	   fmt.Println("Value is 20")
	case 30:
	   fmt.Println("Value is 30")
	case 40:
	   fmt.Println("Value is 40")
	default:
	   fmt.Println("Value is not present in 10 20 30 40")
	}
}