package offices

import (
	"backend/businesses/offices"
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
	Rate        string         `json:"rate"`
}

func FromDomain(domain *offices.Domain) *Office {
	return &Office{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Title:       domain.Title,
		Description: domain.Description,
		City:        domain.City,
		Rate:        domain.Rate,
	}
}

func (rec *Office) ToDomain() offices.Domain {
	return offices.Domain{
		ID:          rec.ID,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
		Title:       rec.Title,
		Description: rec.Description,
		City:        rec.City,
		Rate:        rec.Rate,
	}
}
