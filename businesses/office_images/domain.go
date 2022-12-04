package officeimages

type Domain struct {
	ID uint
	URL string
	OfficeID uint
}

type Usecase interface {
	GetByOfficeID(id string) []Domain
}

type Repository interface {
	GetByOfficeID(id string) []Domain
}