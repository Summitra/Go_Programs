// Program to compare string using Compare()

package main
import (
  "fmt"
  "strings"
)
func main() {
// create three strings
  s1 := "Programiz"
  s2 := "Programiz Pro"
  s3 := "Programiz"

// compare strings
  fmt.Println(strings.Compare(s1, s2))  // -1
  fmt.Println(strings.Compare(s2, s3))  // 1
  fmt.Println(strings.Compare(s1, s3))  // 0
}