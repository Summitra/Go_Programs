package main
import "fmt"
func main(){
	
	var value interface{}=20
	var value1 int = value.(int)

	fmt.Println(value1)

	value2, ok := value.(int)
	if ok{

	   fmt.Println("String value found!!")
	   fmt.Println(value2)
	}else{
	    fmt.Println("String value not found !!")
	}
}