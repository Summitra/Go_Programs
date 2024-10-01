package main
import "fmt"
func mult(a int)(int, float64){
	return a/2, float64(a)*2.05
}
func main(){
	x , y :=mult(12)
	fmt.Println(x,y)
}