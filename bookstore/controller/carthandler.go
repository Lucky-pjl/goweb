package controller

import (
	"encoding/json"
	"goweb/bookstore/dao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"net/http"
	"strconv"
	"text/template"
)

// 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	// 判断是否登录
	flag, session := dao.IsLogin(r)
	if !flag {
		w.Write([]byte("请先登录!"))
		return
	}
	bookId := r.FormValue("bookId")
	// 根据id获取图书信息
	book := dao.GetBookById(bookId)
	// 判读数据库中是否有当前用户的购物车
	userId := session.UserID
	cart, _ := dao.GetCartByUserID(userId)
	if cart != nil {
		// 判断购物车中是否有当前图书
		cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookId, cart.CartID)
		if cartItem != nil {
			cts := cart.CartItems
			for _, v := range cts {
				if cartItem.Book.ID == v.Book.ID {
					v.Count = v.Count + 1
					dao.UpdateBookCount(v)
				}
			}
		} else {
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cart.CartID,
			}
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCartItem(cartItem)
		}
		dao.UpdateCart(cart)
	} else {
		// 1.创建购物车
		cartId := utils.CreateUUID()
		cart := &model.Cart{
			CartID: cartId,
			UserID: userId,
		}
		// 2.创建购物项
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: cartId,
		}
		cartItems = append(cartItems, cartItem)
		cart.CartItems = cartItems
		dao.AddCart(cart)
	}
	w.Write([]byte("将<<" + book.Title + ">>购物车!"))
}

// 获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userId := session.UserID
	cart, _ := dao.GetCartByUserID(userId)
	if cart != nil {
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	// _, session := dao.IsLogin(r)
	// userId := session.UserID
	cartId := r.FormValue("cartId")
	dao.DeleteCartByCartID(cartId)
	GetCartInfo(w, r)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	_, session := dao.IsLogin(r)
	userId := session.UserID
	cart, _ := dao.GetCartByUserID(userId)
	// 获取购物车中的购物项
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartItemId {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			dao.DeleteCartItemByID(cartItemId)
		}
	}
	dao.UpdateCart(cart)
	GetCartInfo(w, r)
}

//UpdateCartItem 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.IsLogin(r)
	userID := session.UserID
	//获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		//寻找要更新的购物项
		if v.CartItemID == iCartItemID {
			v.Count = iBookCount
			dao.UpdateBookCount(v)
		}
	}
	//更新购物车中的图书的总数量和总金额
	dao.UpdateCart(cart)
	cart, _ = dao.GetCartByUserID(userID)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemID {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//将data转换为json字符串
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
