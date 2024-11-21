package main

import (
	"github.com/3169a7e4c0eb100/go-transactional-func/database"
	"github.com/3169a7e4c0eb100/go-transactional-func/payment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PointWallet struct {
	ID      int
	UserID  int
	Balance int
}

func (PointWallet) TableName() string {
	return "point_wallet"
}

type PointTransaction struct {
	ID     int
	UserID int
	Point  int
}

func (PointTransaction) TableName() string {
	return "point_transaction"
}

func main() {

	dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
	}

	pm := payment.NewPaymentService(database.NewGormPointTransactionalDatabase(db))
	err = pm.Pay(2, 100)
	if err != nil {
		panic(err)
	}
}
