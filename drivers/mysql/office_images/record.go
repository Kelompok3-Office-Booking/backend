package officeimages

import (
	officeImageUseCase "backend/businesses/office_images"
	"backend/drivers/mysql/offices"
)

type OfficeImage struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	URL      string `json:"url"`
	Office   offices.Office `json:"office"`
	OfficeID uint   `json:"office_id"`
}

func FromDomain(domain *officeImageUseCase.Domain) *OfficeImage{
	return &OfficeImage{
		ID: domain.ID,
		URL: domain.URL,
		OfficeID: domain.OfficeID,
	}
}

func (rec *OfficeImage) ToDomain() officeImageUseCase.Domain {
	return officeImageUseCase.Domain{
		ID: rec.ID,
		URL: rec.URL,
		OfficeID: rec.OfficeID,
	}
}