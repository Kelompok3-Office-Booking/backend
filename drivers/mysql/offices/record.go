package offices

import (
	officeUsecase "backend/businesses/offices"
	"time"

	"gorm.io/gorm"
)

type Office struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	City        string         `json:"city"`
	Rate        uint           `json:"rate"`
}

func FromDomain(domain *officeUsecase.Domain) *Office {
	return &Office{
		ID:          domain.ID,
		Title:       domain.Title,
		Description: domain.Description,
		City:        domain.City,
		Rate:        domain.Rate,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}

func (rec *Office) ToDomain() officeUsecase.Domain {
	return officeUsecase.Domain{
		ID:          rec.ID,
		Title:       rec.Title,
		Description: rec.Description,
		City:        rec.City,
		Rate:        rec.Rate,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}
