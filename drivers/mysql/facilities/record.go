package facilities

import (
	facilityUsecase "backend/businesses/facilities"
)

type Facility struct {
	ID          uint   `gorm:"primaryKey" json:"id" form:"id"`
	Description string `json:"description" form:"description"`
}

func FromDomain(domain *facilityUsecase.Domain) *Facility {
	return &Facility{
		ID:          domain.ID,
		Description: domain.Description,
	}
}

func (rec *Facility) ToDomain() facilityUsecase.Domain {
	return facilityUsecase.Domain{
		ID:          rec.ID,
		Description: rec.Description,
	}
}
