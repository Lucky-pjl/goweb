package model

import (
	"fmt"
	"goweb/demo/db/utils"
)

// User结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (user *User) String() string {
	return fmt.Sprintln("ID:", user.ID, ",Username:", user.Username,
		",Password:", user.Password, ",Email:", user.Email)
}

// AddUser
func (user *User) AddUser() error {
	// 1.写sql语句
	sqlStr := "insert into users(username,password,email) values (?,?,?)"
	// 2.预编译
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常,err =", err)
		return err
	}
	// 3.执行
	_, err2 := stmt.Exec(user.Username, user.Password, user.Email)
	if err2 != nil {
		fmt.Println("sql执行异常,err =", err)
		return err2
	}
	return nil
}

// 获取单个用户
func (user *User) GetUserById() (u *User, err error) {
	sqlStr := "select * from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, user.ID)
	var id int
	var username, password, email string
	err = row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u = &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return
}

// 获取全部用户
func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select * from users"
	row, err := utils.Db.Query(sqlStr)

	if err != nil {
		return nil, err
	}
	// 创建User切片
	var users []*User
	for row.Next() {
		var id int
		var username, password, email string
		err = row.Scan(&id, &username, &password, &email)
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, err
}
