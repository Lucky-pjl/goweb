package model

import "fmt"

type Book struct {
	ID      int
	Title   string
	Author  string
	Price   float64
	Sales   int
	Stock   int
	ImgPath string
}

func (book *Book) String() string {
	return fmt.Sprintln("ID:", book.ID, ",Title:", book.Title,
		",Author:", book.Author, ",Price:", book.Price, ",Sales:", book.Sales,
		",Stock:", book.Stock, "ImgPath:", book.ImgPath)
}
