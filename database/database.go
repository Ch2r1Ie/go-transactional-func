package database

import (
	"github.com/3169a7e4c0eb100/go-transactional-func/payment"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormPointTransactionalDatabase struct {
	db *gorm.DB
}

func NewGormPointTransactionalDatabase(db *gorm.DB) *gormPointTransactionalDatabase {
	return &gormPointTransactionalDatabase{db: db}
}

func (db *gormPointTransactionalDatabase) Transactional(fn func(tx *gorm.DB) error) error {
	return db.db.Debug().Transaction(fn)
}

func (db *gormPointTransactionalDatabase) InsertPointTransaction(tx *gorm.DB, transaction payment.PointTransaction) error {
	return tx.Debug().Create(&transaction).Error
}

func (db *gormPointTransactionalDatabase) UpsertPointBalance(tx *gorm.DB, userID int, point int) error {
	// Start a transaction
	err := tx.Debug().Transaction(func(tx *gorm.DB) error {
		// Acquire a row-level lock
		var wallet payment.PointWallet
		err := tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id = ?", userID).First(&wallet).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// If no rows were found, insert a new record
				return tx.Debug().Create(&payment.PointWallet{
					UserID:  userID,
					Balance: point,
				}).Error
			}
			return err
		}

		err = tx.Debug().Model(&payment.PointWallet{}).Where("user_id = ?", userID).Update("balance", gorm.Expr("balance + ?", point)).Error
		// If a row was found, update the balance
		return err
	})

	return err
}
