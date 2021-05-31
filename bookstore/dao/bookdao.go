package dao

import (
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"math"
	"strconv"
)

// 获取所有图书
func GetBooks() ([]*model.Book, error) {
	sql := "select * from books"
	row, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for row.Next() {
		var book *model.Book = &model.Book{}
		row.Scan(&book.ID, &book.Title, &book.Author, &book.Price,
			&book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, err
}

// 添加一本图书
func AddBook(book *model.Book) error {
	sql := "insert into books(title,author,price,sales,stock,img_path)values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	return err
}

// 删除一本图书
func DelBook(bookId string) error {
	sql := "delete from books where id = ?"
	_, err := utils.Db.Exec(sql, bookId)
	return err
}

// 获取一本图书
func GetBookById(bookId string) *model.Book {
	sql := "select * from books where id = ?"
	row := utils.Db.QueryRow(sql, bookId)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price,
		&book.Sales, &book.Stock, &book.ImgPath)
	return book
}

// 更新一本图书
func UpdateBook(book *model.Book) error {
	sql := "update books set title = ?,author = ?,price = ?,sales = ?,stock = ? where id = ?"
	_, err := utils.Db.Exec(sql, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	// fmt.Println(err)
	return err
}

func GetPageBooks(pageNo string) (*model.Page, error) {
	sql := "select count(*) from books"
	var totalRecord int64
	row := utils.Db.QueryRow(sql)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	tmp := math.Floor(float64(totalRecord) / float64(pageSize))

	// 获取当前页的图书
	sql2 := "select * from books limit ?,?"
	no, _ := strconv.ParseInt(pageNo, 10, 64)
	row2, err := utils.Db.Query(sql2, (no-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for row2.Next() {
		var book *model.Book = &model.Book{}
		row2.Scan(&book.ID, &book.Title, &book.Author, &book.Price,
			&book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      no,
		PageSize:    pageSize,
		TotalPageNo: int64(tmp),
		TotalRecord: totalRecord,
	}
	return page, err
}

// 获取指定价格区间内的图书
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	sql := "select count(*) from books where price between ? and ?"
	var totalRecord int64
	row := utils.Db.QueryRow(sql, minPrice, maxPrice)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	tmp := math.Floor(float64(totalRecord) / float64(pageSize))

	// 获取当前页的图书
	sql2 := "select * from books where price between ? and ? limit ?,?"
	no, _ := strconv.ParseInt(pageNo, 10, 64)
	row2, err := utils.Db.Query(sql2, minPrice, maxPrice, (no-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for row2.Next() {
		var book *model.Book = &model.Book{}
		row2.Scan(&book.ID, &book.Title, &book.Author, &book.Price,
			&book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	page := &model.Page{
		Books:       books,
		PageNo:      no,
		PageSize:    pageSize,
		TotalPageNo: int64(tmp),
		TotalRecord: totalRecord,
	}
	return page, err
}
