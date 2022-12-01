package request

import (
	"backend/businesses/offices"
	"time"

	"github.com/go-playground/validator/v10"
)

type Office struct {
	Title        string    `json:"title" form:"title" validate:"required"`
	Description  string    `json:"description" form:"description" validate:"required"`
	OfficeType   string    `json:"office_type" form:"office_type" validate:"required,oneof='office' 'coworking' 'meeting room'"`
	OfficeLength uint      `json:"office_length" form:"office_length" validate:"required"`
	PricePerHour uint      `json:"price_per_hour" form:"price_per_hour" validate:"required"`
	OpenHour     time.Time `json:"open_hour" form:"open_hour" validate:"required"`
	CloseHour    time.Time `json:"close_hour" form:"close_hour" validate:"required"`
	Lat          float64   `json:"lat" form:"lat" validate:"required"`
	Lng          float64   `json:"lng" form:"lng" validate:"required"`
	Accommodate  uint      `json:"accommodate" form:"accommodate" validate:"required"`
	WorkingDesk  uint      `json:"working_desk" form:"working_desk" validate:"required"`
	MeetingRoom  uint      `json:"meeting_room" form:"meeting_room" validate:"required"`
	PrivateRoom  uint      `json:"private_room" form:"private_room" validate:"required"`
	City         string    `json:"city" form:"city" validate:"required,oneof='central jakarta' 'south jakarta' 'west jakarta' 'east jakarta' 'thousand islands'"`
	District     string    `json:"district" form:"district" validate:"required"`
	Address      string    `json:"address" form:"address" validate:"required"`
	Rate         float64   `json:"rate"`
}

type OfficeTemp struct {
	Title        string  `json:"title" form:"title" validate:"required"`
	Description  string  `json:"description" form:"description" validate:"required"`
	OfficeType   string  `json:"office_type" form:"office_type" validate:"required,oneof='office' 'coworking' 'meeting room'"`
	OfficeLength uint    `json:"office_length" form:"office_length" validate:"required"`
	PricePerHour uint    `json:"price_per_hour" form:"price_per_hour" validate:"required"`
	OpenHour     string  `json:"open_hour" form:"open_hour" validate:"required"`
	CloseHour    string  `json:"close_hour" form:"close_hour" validate:"required"`
	Lat          float64 `json:"lat" form:"lat" validate:"required"`
	Lng          float64 `json:"lng" form:"lng" validate:"required"`
	Accommodate  uint    `json:"accommodate" form:"accommodate" validate:"required"`
	WorkingDesk  uint    `json:"working_desk" form:"working_desk" validate:"required"`
	MeetingRoom  uint    `json:"meeting_room" form:"meeting_room" validate:"required"`
	PrivateRoom  uint    `json:"private_room" form:"private_room" validate:"required"`
	City         string  `json:"city" form:"city" validate:"required,oneof='central jakarta' 'south jakarta' 'west jakarta' 'east jakarta' 'thousand islands'"`
	District     string  `json:"district" form:"district" validate:"required"`
	Address      string  `json:"address" form:"address" validate:"required"`
	Rate         float64 `json:"rate"`
}

func (req *Office) ToDomain() *offices.Domain {
	return &offices.Domain{
		Title:        req.Title,
		Description:  req.Description,
		OfficeType:   req.OfficeType,
		OfficeLength: req.OfficeLength,
		PricePerHour: req.PricePerHour,
		OpenHour:     req.OpenHour,
		CloseHour:    req.CloseHour,
		Lat:          req.Lat,
		Lng:          req.Lng,
		Accommodate:  req.Accommodate,
		WorkingDesk:  req.WorkingDesk,
		MeetingRoom:  req.MeetingRoom,
		PrivateRoom:  req.PrivateRoom,
		City:         req.City,
		District:     req.District,
		Address:      req.Address,
		Rate:         req.Rate,
	}
}

func (req *Office) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
