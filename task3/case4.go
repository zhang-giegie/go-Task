package task3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Book 书籍结构体
type Book struct {
	ID     int     `db:"id" json:"id"`
	Title  string  `db:"title" json:"title"`
	Author string  `db:"author" json:"author"`
	Price  float64 `db:"price" json:"price"`
}

func GetExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	const query = `
		SELECT id, title, author, price 
		FROM books 
		WHERE price > ? 
		ORDER BY price DESC
	`

	var books []Book
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询高价书籍失败: %v", err)
	}

	return books, nil
}

// Run 执行示例
func RunCase4(db *sqlx.DB) {
	//创建表，并模拟数据
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			price REAL NOT NULL
		)
	`)
	db.MustExec(`
		INSERT INTO books (title, author, price) VALUES 
		('《Go 语言基础》', ' effective-go', 39.9),
		('《Go 语言实战》', ' effective-go', 59.9),
		('《Go 语言微服务》', ' effective-go', 49.9),
		('《Go 语言并发编程》', ' effective-go', 69.9)
	`)

	// 查询价格大于50元的书籍
	books, err := GetExpensiveBooks(db, 50.0)
	if err != nil {
		fmt.Printf("查询失败: %v\n", err)
		return
	}

	fmt.Printf("找到%d本价格超过50元的书籍:\n", len(books))
	for _, book := range books {
		fmt.Printf("- %s (作者: %s, 价格: ￥%.2f)\n",
			book.Title, book.Author, book.Price)
	}
}
