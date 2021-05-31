package controller

import (
	"goweb/bookstore/dao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"net/http"
	"text/template"
	"time"
)

// 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userId := session.UserID
	cart, _ := dao.GetCartByUserID(userId)
	// 创建Order
	orderId := utils.CreateUUID()
	order := &model.Order{
		OrderID:     orderId,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userId),
	}
	dao.AddOrder(order)
	// 保存订单项
	cartItems := cart.CartItems
	for _, v := range cartItems {
		// 创建OrderItem
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderId,
		}
		dao.AddOrderItem(orderItem)
		// 更新当前购物项中图书的库存和销量
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		dao.UpdateBook(book)
	}

	// 情空购物车
	dao.DeleteCartByCartID(cart.CartID)
	session.OrderID = orderId
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

// 获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderID(orderId)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

// 获取订单
func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userId := session.UserID
	orders, _ := dao.GetMyOrders(userId)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

// 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, 1)
	GetOrders(w, r)
}

// 收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	dao.UpdateOrderState(orderId, 2)
	GetMyOrder(w, r)
}
