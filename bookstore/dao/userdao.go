package dao

import (
	"goweb/bookstore/model"
	"goweb/bookstore/utils"
)

// 验证账户和密码,从数据库查询记录
func CheckUser(username string, password string) (user *model.User, err error) {
	sql := "select * from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sql, username, password)
	user = &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return
}

// 根据用户名查询
func CheckUsername(username string) (user *model.User, err error) {
	sql := "select * from users where username=?"
	row := utils.Db.QueryRow(sql, username)
	user = &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return
}

// SaveUser
func SaveUser(username string, password string, email string) error {
	sql := "insert into users(username,password,email) values (?,?,?)"
	_, err := utils.Db.Exec(sql, username, password, email)
	return err
}
