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
	OfficeType 	string
	OfficeLength uint
	PricePerHour uint
	Lat			float64
	Lng			float64
	Accommodate	uint
	WorkingDesk	uint
	MeetingRoom	uint
	PrivateRoom uint
	City        string
	District	string
	Address		string
	Rate        float64
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
