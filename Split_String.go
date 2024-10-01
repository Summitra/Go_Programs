package main
import (
  "fmt"
  "strings"
)

func main() {
  var message = "I Love Golang"
  
  // split string from space " "
  splittedString := strings.Split(message, " ")


  fmt.Println(splittedString)
}

// Output: [I Love Golang]

//The Split() method returns a slice of all the substrings. In our example, [I Love Golang] is a slice.
