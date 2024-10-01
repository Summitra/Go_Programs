package main
import "fmt"
func mult(a int)(int, int){
	return a/2 , a*2
}
func main(){
	x,y :=mult(12)
	fmt.Println(x,y)
}