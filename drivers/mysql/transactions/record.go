package transactions

import (
	TransactionUseCase "backend/businesses/transactions"
	"backend/drivers/mysql/offices"
	"backend/drivers/mysql/users"
)

type Transaction struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Price    uint           `json:"price"`
	UserID   uint           `json:"user_id"`
	OfficeID uint           `json:"office_id"`
	User     users.User     `json:"user" gorm:""`
	Office   offices.Office `json:"office" gorm:""`
}

func FromDomain(domain *TransactionUseCase.Domain) *Transaction {
	return &Transaction{
		ID:       domain.ID,
		Price:    domain.Price,
		UserID:   domain.UserID,
		OfficeID: domain.OfficeID,
	}
}

func (rec *Transaction) ToDomain() TransactionUseCase.Domain {
	return TransactionUseCase.Domain{
		ID:       rec.ID,
		Price:    rec.Price,
		UserID:   rec.UserID,
		OfficeID: rec.OfficeID,
	}
}
