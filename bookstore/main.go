package main

import (
	"goweb/bookstore/controller"
	"net/http"
)

func main() {
	// 设置处理静态资源
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/",
		http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	http.HandleFunc("/main", controller.IndexHandler)
	// 登录
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	// 注册
	http.HandleFunc("/regist", controller.Regist)
	// 通过Ajax验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUsername)
	// 获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	http.HandleFunc("/addBook", controller.UpdateBook)
	http.HandleFunc("/delBook", controller.DelBook)
	http.HandleFunc("/toUpdateBook", controller.ToUpdateBook)
	http.HandleFunc("/updateBook", controller.UpdateBook)
	// 购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	// 订单
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)
	// 创建路由
	http.ListenAndServe(":8080", nil)
}
