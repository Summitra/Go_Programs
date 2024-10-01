package main
import "fmt"
func main(){
   var n int
   fmt.Print("Enter any number: ")
   fmt.Scanf("%d", &n)
   s:= 0
   for i:=1; i<n;i++{
      if n % i ==0{
         s += i
      }
   }
   if s == n {
      print("The number is a Perfect number!")
   }else{
      print("The number is not a Perfect number!")
   }
}