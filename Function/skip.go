package main
import "fmt"
func abc (a int)(int,int){
	return a*2, a*4
}
func main(){
	x, _:=abc(2)
	fmt.Println(x)
}