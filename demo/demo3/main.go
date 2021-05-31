package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求地址:", r.URL.Path)
	fmt.Fprintln(w, "查询字符串:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头信息:", r.Header["Accept-Encoding"])
	// 获取请求体内容的长度
	len := r.ContentLength
	// 创建byte切片
	body := make([]byte, len)
	// 读取请求体内容
	r.Body.Read(body)
	fmt.Fprintln(w, "请求体中的内容:", string(body))

	// 解析表单
	r.ParseForm()
	fmt.Fprintln(w, "请求参数有:", r.Form)
	fmt.Fprintln(w, "请求参数有:", r.PostForm)
}

// 响应josn数据
func handler2(w http.ResponseWriter, r *http.Request) {
	// 设置响应内容类型
	w.Header().Set("Content-Type", "application/json")
	user := User{
		ID:       1,
		Username: "张三",
		Password: "123456",
		Email:    "zs@qq.com",
	}
	json, _ := json.Marshal(user)
	w.Write(json)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/res", handler2)
	// 创建路由
	http.ListenAndServe(":8080", nil)
}
