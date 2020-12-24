package models

import (
	"Item/dbmysql"
	"Item/utils"
	"log"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

//用户信息保存到数据库
func (u User) AddUser() (int64, error) {
	u.Password = utils.MD5HashString(u.Password)
	rs, err := dbmysql.Db.Exec("insert into user(phone,password) values(?,?)",
		u.Phone, u.Password)
	if err != nil { //保存数据遇到错误
		//fmt.Println(err.Error())
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//查询信息
func (u User) QuerUser() (*User, error) {
	u.Password = utils.MD5HashString(u.Password)
	/*row := dbmysql.Db.QueryRow("select phone,password from user where phone = ? and password = ?",
	u.Phone, u.Password)
	*/
	row, err := dbmysql.Db.Query("select phone,password from user where phone = ? and password = ?",
		u.Phone, u.Password)
	if err != nil {
		log.Fatal(err)
	}
	//defer row.Close()
	for row.Next() {
		err := row.Scan(&u.Phone, &u.Password)
		if err != nil {
			log.Fatal(err)
			return nil,err
		}
		log.Fatal(u.Phone, u.Password)
	}
	return &u, nil
}
