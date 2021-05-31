package main

import (
	"html/template"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	// 1.解析模板
	// t, _ := template.ParseFiles("index.html")
	// Must函数处理异常
	t := template.Must(template.ParseFiles("index.html"))
	// 2.执行
	t.Execute(w, "Hello Template")
}

func main() {
	http.HandleFunc("/template", handler)
	// 创建路由
	http.ListenAndServe(":8080", nil)
}
