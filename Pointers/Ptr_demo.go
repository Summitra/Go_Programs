package main
import "fmt"
func main(){
	var a int =20 // actual value
	var p *int
	p= &a
	fmt.Printf("Address of varialbe %x", &a)
	fmt.Printf("\nAddress of variable %x",p)
	fmt.Printf("\nValue of variable %d", *p)
	fmt.Printf("\nValue of variable %d", a)
}