package main

import(
	"fmt"
)

type Book struct{
	Title string
	Name string
	auth string

}

func changeBook(book *Book){
	book.Title = book.Title + "-New"
}

func changeBook2(book Book){
	book.Title = book.Title + "-Old"
	fmt.Println(book)
}

func main(){
	var book1 = Book{Title:"Golang",Name:"gogogo",auth:"paul"}
	fmt.Println(book1)

	changeBook(&book1)
	fmt.Println(book1)

	changeBook2(book1)
	fmt.Println(book1)
}