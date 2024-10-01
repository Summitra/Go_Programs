// Program using Replace() to replace strings

package main
import (
  "fmt"
  "strings"
)

func main() {
    
  text := "car"
  fmt.Println("Old String:", text)
  
  // replace r with t                                       
  replacedText := strings.Replace(text, "r", "t", 1)   //r is old character to replace
   fmt.Println("New String:", replacedText)            // t is new character to replace
  }                                                     // 1 how many old characters to replace
       

       // replace 2 r with 2 a
       //  strings.Replace("Programiz", "r", "R", 2)
       // Output: PRogRamiz
