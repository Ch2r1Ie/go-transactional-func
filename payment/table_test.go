package payment

import "testing"

func TestTableNamePointWallet(t *testing.T) {
	table := PointWallet{}
	if table.TableName() != "point_wallet" {
		t.Error("PointWallet.TableName() should return point_wallet")
	}
}

func TestTableNamePointTransaction(t *testing.T) {
	table := PointTransaction{}
	if table.TableName() != "point_transaction" {
		t.Error("PointTransaction.TableName() should return point_transaction")
	}
}
