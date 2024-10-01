package main
import "fmt"
type Shape interface{
	Sarea() int
	Sperimeter() int
}
type Circle struct{
	radius int
}
type Rectangle struct{
	length, breadth, width int
}
func (c Circle) Sarea() int{
	return 3 * c.radius*c.radius
}
func (c Circle) Sperimeter() int{
	return 2 * 3 * c.radius
}
func (r Rectangle) Sarea() int{
	return r.length * r.breadth
}
func (r Rectangle) Sperimeter()int{
	return 2*(r.length + r.width)
}
func main(){
	var c int 
	var r int 
	c = Circle(12)
	r = Rectangle(10,20)

	fmt.Println("Area of Circle :", c.Sarea())
	fmt.Println("Perimeter of Circle :", c.Sperimeter())

	fmt.Println("Area of Rectanlge :", r.Sarea())
	fmt.Println("Perimeter of Rectangle :" ,r.Sperimeter())
}