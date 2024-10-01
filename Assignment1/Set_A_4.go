package main
import "fmt"

func main(){
// variable
i := 32
fmt.Println("Address of i is : ",&i)

// pointer to the variable
p := &i
fmt.Println("Value of p is : ",p)
}
