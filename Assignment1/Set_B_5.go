
package main
import "fmt"
 
func area(length, width int)int{
     
    Ar := length* width
    return Ar
}
 
// Main function
func main() {
   
     fmt.Printf("Area of rectangle is : %d", area(12, 10))
}