package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Price         uint
	CheckIn       time.Time
	CheckOut      time.Time
	Duration      int
	PaymentMethod string
	Status        string
	Drink         string
	UserID        uint
	OfficeID      uint
}

type Usecase interface {
	GetAll() []Domain
	Create(transactionDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, transactionDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	Create(transactionDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, transactionDomain *Domain) Domain
	Delete(id string) bool
}
