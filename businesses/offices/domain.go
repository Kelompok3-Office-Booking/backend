package offices

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Title       string
	Description string
	City        string
	Rate        uint
}

type OfficeRepository interface {
	Create(officeDomain *Domain) Domain
	GetAll() []Domain
	GetByID(id string) Domain
	Delete(id string) bool
	SearchByCity(city string) []Domain
	SearchByRate(rate string) []Domain
}
