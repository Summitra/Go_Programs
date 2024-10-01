package main
import "fmt"
func show( a int) int {
	x=a+8
	y=a*8
	return
}
func main(){
	_,y := show(3)
	fmt.Println(y)
}