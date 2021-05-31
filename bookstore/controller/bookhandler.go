package controller

import (
	"goweb/bookstore/dao"
	"goweb/bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

// 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	// 1.解析模板
	t, _ := template.ParseFiles("views/index.html")
	// 2.执行
	t.Execute(w, page)
}

// 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	// 解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, books)
}

// 获取所有图书
func DelBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("id")
	dao.DelBook(bookId)
	GetPageBooks(w, r)
}

// 去更新页面
func ToUpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("id")
	book := dao.GetBookById(bookId)
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

// 更新图书或添加图书
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	id, _ := strconv.Atoi(bookId)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.Atoi(sales)
	iStock, _ := strconv.Atoi(stock)
	book := &model.Book{
		ID:      id,
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: "static/img/default.jpg",
	}
	// fmt.Println(book)
	if book.ID > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}

// 获取分页数据
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	// 解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, page)
}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" || maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	// 获取Cookie
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, page)
}
