package transactions

type Domain struct {
	ID       uint
	Price    uint
	UserID   uint
	OfficeID uint
}

type Usecase interface {
	GetAll() []Domain
	Create(transactionDomain *Domain) Domain
}

type Repository interface {
	GetAll() []Domain
	Create(transactionDomain *Domain) Domain
}
