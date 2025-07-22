package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhang-giegie/go-Task/task3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	return db
}

func initSqlxDB() *sqlx.DB {
	// 初始化sqlx连接
	db, err := sqlx.Connect("mysql",
		"root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	db := InitDB()
	//task3.Run(db)
	task3.RunCase5(db)

	//db := initSqlxDB()
	//task3.Run_case3(db)
	//task3.RunCase4(db)
}
