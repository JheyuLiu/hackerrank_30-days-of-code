package main

import (
	"fmt"
	"bufio"
	"os"
)

type Book interface {
	display()
}

type BookInfo struct {
	title string
	author string
}

func (bi BookInfo) display() {
	fmt.Println("Title:", bi.title)
	fmt.Println("Author:", bi.author)
}

type MyBook struct {
	BookInfo

	price int
}

func (b MyBook) display() {
    b.BookInfo.display()
    fmt.Println("Price:", b.price)
}

func main() {
	var price int
    scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
    title := scanner.Text()
    
    scanner.Scan()
	author := scanner.Text()
    
    fmt.Scan(&price)	

	myBook := MyBook{BookInfo{title, author}, price}
	myBook.display()
}