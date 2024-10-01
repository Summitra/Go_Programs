package main

import "fmt"

func main() {

    var n, fact int
    fact = 1

    fmt.Print("Enter any Number to find the Factorial :")
    fmt.Scanln(&n)

    for i := 1; i <= n; i++ {
        fact = fact* i
    }
    fmt.Println("The Factorial of ", n, "is  = ", fact)
}