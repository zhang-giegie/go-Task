package task3

import (
	"fmt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Name  string `gorm:"size:100"`
	Posts []Post // 一对多关系
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"size:200"`
	UserID   uint      // 外键
	Comments []Comment // 一对多关系
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	PostID  uint   // 外键
}

// 查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// 查询评论数量最多的文章信息
func getMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	err := db.
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post).Error
	return post, err
}

// Post 模型添加钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量统计
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).Error
}

// Comment 模型添加钩子函数
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 检查文章的评论数量
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	// 如果评论数量为0，更新文章状态
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("status", "无评论").Error
	}
	return nil
}

func RunCase5(db *gorm.DB) {
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	////生成数据 完整的User包含Post和Comment
	//db.Create(&User{Name: "Alice", Posts: []Post{{Title: "Post 1", Comments: []Comment{{Content: "Comment 1"}, {Content: "Comment 2"}}}, {Title: "Post 2", Comments: []Comment{{Content: "Comment 3"}}}}})
	//db.Create(&User{Name: "Bob", Posts: []Post{{Title: "Post 3", Comments: []Comment{{Content: "Comment 4"}, {Content: "Comment 5"}}}, {Title: "Post 4", Comments: []Comment{{Content: "Comment 6"}}}}})
	//db.Create(&User{Name: "Charlie", Posts: []Post{{Title: "Post 7", Comments: []Comment{{Content: "Comment 7"}, {Content: "Comment 8"}}}, {Title: "Post 8", Comments: []Comment{{Content: "Comment 9"}}}}})
	//db.Create(&User{Name: "David", Posts: []Post{{Title: "Post 9", Comments: []Comment{{Content: "Comment 10"}, {Content: "Comment 11"}, {Content: "Comment 12"}}}, {Title: "Post 10", Comments: []Comment{{Content: "Comment 12"}}}}})

	//result, _ := getMostCommentedPost(db)
	//fmt.Println(result)

	comments, _ := getUserPostsWithComments(db, 1)
	fmt.Println(comments)
}
