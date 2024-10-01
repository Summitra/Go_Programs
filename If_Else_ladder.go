package main
import "fmt"
func main(){
	var a int
	fmt.Print("Enter Marks:")
	fmt.Scanln(&a)
	if(a<0 || a>100){
	   fmt.Printf("Please enter the valid no")
	}else if(a>=0 && a<50){
	   fmt.Printf("Fail")
	}else if(a>=50 && a<60){
	   fmt.Printf("D Grade")
	}else if(a>=60 && a<70){
	   fmt.Printf("C Grade")
	}else if(a>=70 && a<80){
	   fmt.Printf("B Grade")
	}else if(a>=80 && a<90){
	   fmt.Printf("A Grade")
	}else if(a>=90 && a<100){
	   fmt.Printf("A++ Grade")
	}
}