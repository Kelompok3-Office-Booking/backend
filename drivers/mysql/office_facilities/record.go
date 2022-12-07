package officefacilities

import (
	officeFacilityUseCase "backend/businesses/office_facilities"
	"backend/drivers/mysql/offices"
)

type OfficeFacility struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FacilitiesID string         `json:"facilities_id"`
	OfficeID     uint           `json:"office_id"`
	Office       offices.Office `json:"office" gorm:""`
}

func FromDomain(domain *officeFacilityUseCase.Domain) *OfficeFacility {
	return &OfficeFacility{
		ID:           domain.ID,
		FacilitiesID: domain.FacilitiesID,
		OfficeID:     domain.OfficeID,
	}
}

func (rec *OfficeFacility) ToDomain() officeFacilityUseCase.Domain {
	return officeFacilityUseCase.Domain{
		ID:           rec.ID,
		FacilitiesID: rec.FacilitiesID,
		OfficeID:     rec.OfficeID,
	}
}
