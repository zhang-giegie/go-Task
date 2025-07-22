package task3

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 `gorm:"not null"`
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountID uint    `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	CreatedAt     time.Time
}

func TransferMoneyGORM(db *gorm.DB, fromID, toID uint, amount float64) error {
	// 开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 1. 检查转出账户余额
	var fromAccount Account
	if err := tx.First(&fromAccount, fromID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %v", err)
	}

	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足")
	}

	// 2. 扣除转出账户金额
	if err := tx.Model(&Account{}).
		Where("id = ?", fromID).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣除金额失败: %v", err)
	}

	// 3. 增加转入账户金额
	if err := tx.Model(&Account{}).
		Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("增加金额失败: %v", err)
	}

	// 4. 记录交易
	transaction := Transaction{
		FromAccountID: fromID,
		ToAccountID:   toID,
		Amount:        amount,
		CreatedAt:     time.Now(),
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %v", err)
	}

	// 提交事务
	return tx.Commit().Error
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transaction{})

	//创建两个账户
	db.Create(&Account{ID: 1, Balance: 1000.0})
	db.Create(&Account{ID: 2, Balance: 500.0})

	err := TransferMoneyGORM(db, 1, 2, 100.0)
	if err != nil {
		fmt.Printf("转账失败: %v\n", err)
	} else {
		fmt.Println("转账成功")
	}
}
