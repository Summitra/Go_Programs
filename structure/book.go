package main
import "fmt"
type Books struct{
	title string
	author string
	subject string
	book_id int
	price int
}
func main(){
	var b1[10]Books
	var n int
	fmt.Println("Enter books you want:")
	fmt.Scan(&n)
	for i:=0; i<n;i++{
	fmt.Println("Enter Title:")
	fmt.Scan(&b1[i].title)
	fmt.Println("Enter author name:")
	fmt.Scan(&b1[i].author)
	fmt.Println("Enter Subject:")
	fmt.Scan(&b1[i].subject)
	fmt.Println("Enter Book Id:")
	fmt.Scan(&b1[i].book_id)
	fmt.Println("Enter the price:")
	fmt.Scan(&b1[i].price)
	}
	for i:= 0; i<n ;i++ {
	fmt.Println("Book Title :", b1[i]. title)
	fmt.Println("Book Author :" , b1[i].author)
	fmt.Println("Book Subject :", b1[i].subject)
	fmt.Println("Book Id :", b1[i].book_id)
	fmt.Println("Book Price :", b1[i].price)
	fmt.Println("**************")
	}
}