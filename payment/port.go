//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"

package payment

import "gorm.io/gorm"

type PointTransactionalDatabase interface {
	Transactional(fn func(tx *gorm.DB) error) error
	InsertPointTransaction(tx *gorm.DB, pt PointTransaction) error
	UpsertPointBalance(tx *gorm.DB, userID int, point int) error
}
