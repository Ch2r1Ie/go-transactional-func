package payment

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
