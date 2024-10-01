package main
import "fmt"
func sortArray(arr [5]int, min int, temp int) [5]int {
   for i := 0; i <= 4; i++ {
      min = i
      for j := i + 1; j <= 4; j++ {
         if arr[j] < arr[min] {
         
            // changing the index to show the min value
            min = j
         }
      }
      temp = arr[i]
      arr[i] = arr[min]
      arr[min] = temp
   }
   return arr
}
func main() {
   arr := [5]int{50, 30, 20, 10, 40}
   fmt.Println("The unsorted array entered is:", arr)
   var min int = 0
   var temp int = 0
   array := sortArray(arr, min, temp)
   fmt.Println()
   fmt.Println("The Sorted Array is :", array)
}