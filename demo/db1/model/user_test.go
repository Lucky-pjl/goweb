package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("测试查询单个用户:", testGetUserById)
	t.Run("测试查询多条用户:", testGetUsers)
}

func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户: ")
	user := &User{
		Username: "zs",
		Password: "zs",
		Email:    "zs@qq.com",
	}
	user.AddUser()
}

func testGetUserById(t *testing.T) {
	fmt.Println("测试一条记录: ")
	user := User{
		ID: 1,
	}
	u, _ := user.GetUserById()
	fmt.Print(u)
}

func testGetUsers(t *testing.T) {
	fmt.Println("测试多条记录: ")
	user := &User{}
	users, _ := user.GetUsers()
	for _, u := range users {
		fmt.Print(u)
	}
}
