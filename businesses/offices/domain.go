package offices

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Domain struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	City        string         `json:"city"`
	Rate        uint           `json:"rate"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type DomainInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	City        string `json:"city" validate:"required"`
	Rate        uint   `json:"rate" validate:"required"`
}

type OfficeRepository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(input DomainInput) Domain
	Update(id string, input DomainInput) Domain
	Delete(id string) bool
}

type SearchOffice interface {
	SearchByCity(city string) []Domain
	SearchByRate(rate string) []Domain
}

func (input *DomainInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
