package transactions

import (
	transactions "backend/businesses/transactions"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) transactions.Repository {
	return &TransactionRepository{
		conn: conn,
	}
}

func (t *TransactionRepository) GetAll() []transactions.Domain {
	var rec []Transaction

	t.conn.Find(&rec)

	TransactionDomain := []transactions.Domain{}

	for _, trans := range rec {
		TransactionDomain = append(TransactionDomain, trans.ToDomain())
	}

	return TransactionDomain
}

func (t *TransactionRepository) Create(TransactionDomain *transactions.Domain) transactions.Domain {
	rec := FromDomain(TransactionDomain)

	result := t.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}
