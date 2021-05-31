package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	// fmt.Println("测试开始...")
	// t.Run("测试账户:", testLogin)
	// t.Run("测试注册:", testRegist)
	// t.Run("测试添加:", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUser("zs", "zs")
	fmt.Println(user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUsername("ls")
	fmt.Println(user)
}

func testSave(t *testing.T) {
	SaveUser("ls", "ls", "ls@qq.com")
}
