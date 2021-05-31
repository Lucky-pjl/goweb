package dao

import (
	"fmt"
	"goweb/bookstore/model"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestBooks(t *testing.T) {
	fmt.Println("测试开始...")
	t.Run("测试:", testGetBooks)
	t.Run("测试:", testAddBook)
}

func testGetBooks(t *testing.T) {
	fmt.Println("获取所有图书:")
	books, _ := GetBooks()
	for _, book := range books {
		fmt.Println(book)
	}
}

func testAddBook(t *testing.T) {
	fmt.Println("添加图书:")
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.8,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	AddBook(book)
}
