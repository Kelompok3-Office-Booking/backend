package response

import (
	"backend/businesses/offices"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Office struct {
	ID           uint     `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	DeletedAt    string   `json:"deleted_at"`
	Title        string   `json:"title" form:"title"`
	Description  string   `json:"description" form:"description"`
	OfficeType   string   `json:"office_type" form:"office_type"`
	OfficeLength uint     `json:"office_length" form:"office_length"`
	PricePerHour uint     `json:"price_per_hour" form:"price_per_hour"`
	OpenHour     string   `json:"open_hour" form:"open_hour"`
	CloseHour    string   `json:"close_hour" form:"close_hour"`
	Lat          float64  `json:"lat" gorm:"type:decimal(10,7)" form:"lat"`
	Lng          float64  `json:"lng" gorm:"type:decimal(11,7)" form:"lng"`
	Accommodate  uint     `json:"accommodate" form:"accommodate"`
	WorkingDesk  uint     `json:"working_desk" form:"working_desk"`
	MeetingRoom  uint     `json:"meeting_room" form:"meeting_room"`
	PrivateRoom  uint     `json:"private_room" form:"private_room"`
	City         string   `json:"city" form:"city"`
	District     string   `json:"district" form:"district"`
	Address      string   `json:"address" form:"address"`
	Rate         float64  `json:"rate" form:"rate"`
	Images       []string `json:"images" form:"images"`
}

func FromDomain(domain offices.Domain) Office {
	return Office{
		ID:           domain.ID,
		Title:        domain.Title,
		Description:  domain.Description,
		OfficeType:   domain.OfficeType,
		OfficeLength: domain.OfficeLength,
		PricePerHour: domain.PricePerHour,
		OpenHour:     domain.OpenHour.Format("15:04"),
		CloseHour:    domain.CloseHour.Format("15:04"),
		Lat:          domain.Lat,
		Lng:          domain.Lng,
		Accommodate:  domain.Accommodate,
		WorkingDesk:  domain.WorkingDesk,
		MeetingRoom:  domain.MeetingRoom,
		PrivateRoom:  domain.PrivateRoom,
		City:         cases.Title(language.English).String(domain.City),
		District:     cases.Title(language.English).String(domain.District),
		Address:      cases.Title(language.English).String(domain.Address),
		Rate:         domain.Rate,
		Images:       domain.Images,
		CreatedAt:    domain.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:    domain.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:    domain.DeletedAt.Time.Format("01-02-2006 15:04:05"),
	}
}