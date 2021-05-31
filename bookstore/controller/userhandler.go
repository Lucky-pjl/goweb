package controller

import (
	"goweb/bookstore/dao"
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
	"net/http"
	"text/template"
)

// Login
func Login(w http.ResponseWriter, r *http.Request) {

	flag, _ := dao.IsLogin(r)
	if flag {
		GetPageBooksByPrice(w, r)
		return
	}
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	// 调用UserDao中的方法
	user, _ := dao.CheckUser(username, password)
	if user.ID > 0 {
		// 生成uuid
		uuid := utils.CreateUUID()
		sess := &model.Session{
			SessionID: uuid,
			UserName:  user.Username,
			UserID:    user.ID,
		}
		// 将session存入数据库
		dao.AddSession(sess)
		// 创建一个Cookie,让它与Session相关联
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		// 将Cookie发送给浏览器
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		user.Password = ""
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// 获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		// 设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageBooksByPrice(w, r)
}

func Regist(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	// 调用UserDao中的方法
	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		// 用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

// 通过Ajax请求验证用户名是否可用
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	// 获取用户名
	username := r.PostFormValue("username")
	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		// 用户名已存在
		w.Write([]byte("用户名已存在!"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用!</font>"))
	}
}
