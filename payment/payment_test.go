package payment

import (
	"testing"

	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

type PaymentTestSuite struct {
	suite.Suite

	ctrl                           *gomock.Controller
	mockPointTransactionalDatabase *MockPointTransactionalDatabase

	svc *paymentService
}

func TestPaymentTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentTestSuite))
}

func (t *PaymentTestSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())
	t.mockPointTransactionalDatabase = NewMockPointTransactionalDatabase(t.ctrl)

	t.svc = NewPaymentService(t.mockPointTransactionalDatabase)
}

func (t *PaymentTestSuite) TestPayment() {
	userID := 1
	point := 100

	t.mockPointTransactionalDatabase.EXPECT().Transactional(gomock.Any()).DoAndReturn(
		func(fn func(tx *gorm.DB) error) error {
			t.mockPointTransactionalDatabase.EXPECT().InsertPointTransaction(gomock.Any(), PointTransaction{
				UserID: userID,
				Point:  point,
			}).Return(nil)

			t.mockPointTransactionalDatabase.EXPECT().UpsertPointBalance(gomock.Any(), userID, point).Return(nil)

			return fn(nil)
		},
	).Return(nil)

	err := t.svc.Pay(userID, point)
	t.NoError(err)
}
