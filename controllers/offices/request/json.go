package request

import (
	"backend/businesses/offices"

	"github.com/go-playground/validator/v10"
)

type Office struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	City        string `json:"city" validate:"required"`
	Rate        uint   `json:"rate" validate:"required"`
}

func (req *Office) ToDomainCreate() *offices.Domain {
	return &offices.Domain{
		Title:       req.Title,
		Description: req.Description,
		City:        req.City,
	}
}

func (req *Office) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
