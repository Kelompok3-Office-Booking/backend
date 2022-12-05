package response

import facilities "backend/businesses/facilities"

type Facility struct {
	ID          uint   `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
}

func FromDomain(domain facilities.Domain) Facility {
	return Facility{
		ID:          domain.ID,
		Description: domain.Description,
	}
}
