package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//  go get -u github.com/go-sql-driver/mysql

type user struct {
	id   int
	name string
}

var db *sql.DB //连接池对象

func ConnDb() {

	dsn := "root:root@tcp(127.0.0.1:3306)/user"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库成功")
	}
	return
}

func QueryRow() user {

	sqlStr := "select * from user where id=?"

	var u user

	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name)

	if err != nil {
		fmt.Printf("scan db error:%v", err)
	}
	return u
}
