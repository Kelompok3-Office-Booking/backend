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

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(officeDomain *Domain) Domain
	Update(id string, noteDomain *Domain) Domain
	Delete(id string) bool
	SearchByCity(city string) []Domain
	SearchByRate(rate string) []Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(officeDomain *Domain) Domain
	Update(id string, noteDomain *Domain) Domain
	Delete(id string) bool
	SearchByCity(city string) []Domain
	SearchByRate(rate string) []Domain
}
