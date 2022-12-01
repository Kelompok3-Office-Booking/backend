package offices

import (
	officeUsecase "backend/businesses/offices"
	"time"

	"gorm.io/gorm"
)

type Office struct {
	ID           uint           `gorm:"primaryKey" json:"id" form:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Title        string         `json:"title" form:""`
	Description  string         `json:"description" form:"description"`
	OfficeType   string         `gorm:"type:enum('office', 'coworking', 'meeting room')" json:"office_type" form:"office_type"`
	OfficeLength uint           `json:"office_length" form:"office_length"`
	PricePerHour uint           `json:"price_per_hour" form:"price_per_hour"`
	Lat          float64        `gorm:"type:decimal(10,7)" json:"lat" form:"lat"`
	Lng          float64        `gorm:"type:decimal(11,7)" json:"lng" form:"lng"`
	Accommodate  uint           `json:"accommodate" form:"accommodate"`
	WorkingDesk  uint           `json:"working_desk" form:"working_desk"`
	MeetingRoom  uint           `json:"meeting_room" form:"meeting_room"`
	PrivateRoom  uint           `json:"private_room" form:"private_room"`
	City         string         `gorm:"type:enum('central jakarta', 'south jakarta', 'west jakarta', 'east jakarta', 'thousand islands')" json:"city" form:"city"`
	District     string         `json:"district" form:"district"`
	Address      string         `json:"address" form:"address"`
	Rate         float64        `json:"rate" form:"rate"`
}

func FromDomain(domain *officeUsecase.Domain) *Office {
	return &Office{
		ID:           domain.ID,
		Title:        domain.Title,
		Description:  domain.Description,
		OfficeType:   domain.OfficeType,
		OfficeLength: domain.OfficeLength,
		PricePerHour: domain.PricePerHour,
		Lat:          domain.Lat,
		Lng:          domain.Lng,
		Accommodate:  domain.Accommodate,
		WorkingDesk:  domain.WorkingDesk,
		MeetingRoom:  domain.MeetingRoom,
		PrivateRoom:  domain.PrivateRoom,
		City:         domain.City,
		District:     domain.District,
		Address:      domain.Address,
		Rate:         domain.Rate,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}

func (rec *Office) ToDomain() officeUsecase.Domain {
	return officeUsecase.Domain{
		ID:           rec.ID,
		Title:        rec.Title,
		Description:  rec.Description,
		OfficeType:   rec.OfficeType,
		OfficeLength: rec.OfficeLength,
		PricePerHour: rec.PricePerHour,
		Lat:          rec.Lat,
		Lng:          rec.Lng,
		Accommodate:  rec.Accommodate,
		WorkingDesk:  rec.WorkingDesk,
		MeetingRoom:  rec.MeetingRoom,
		PrivateRoom:  rec.PrivateRoom,
		City:         rec.City,
		District:     rec.District,
		Address:      rec.Address,
		Rate:         rec.Rate,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}
