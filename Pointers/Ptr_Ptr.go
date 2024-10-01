package main
import "fmt"
func main(){
	var a= 20
	var p1 = &a
	var p2 = &p1
	fmt.Printf("Value of %d", *p1)
	fmt.Printf("\nValue of a %d", **p2)
}