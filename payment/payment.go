package payment

import (
	"fmt"

	"gorm.io/gorm"
)

type paymentService struct {
	pointTransactionDB PointTransactionalDatabase
}

func NewPaymentService(pointTransactionDB PointTransactionalDatabase) *paymentService {
	return &paymentService{
		pointTransactionDB: pointTransactionDB,
	}
}

func (p *paymentService) Pay(userID, point int) error {
	_ = p.pointTransactionDB.Transactional(func(tx *gorm.DB) error {
		fmt.Println("tx", tx)
		if err := p.pointTransactionDB.InsertPointTransaction(tx, PointTransaction{
			UserID: userID,
			Point:  point,
		}); err != nil {
			return err
		}

		if err := p.pointTransactionDB.UpsertPointBalance(tx, userID, point); err != nil {
			return err
		}

		return nil
	})

	return nil
}
