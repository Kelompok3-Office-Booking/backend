package facilities

import (
	"backend/businesses/facilities"

	ctrl "backend/controllers"
	"backend/controllers/facilities/request"
	"backend/controllers/facilities/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type FacilityController struct {
	facilityUsecase facilities.Usecase
}

func NewFacilityController(facilityUC facilities.Usecase) *FacilityController {
	return &FacilityController{
		facilityUsecase: facilityUC,
	}
}

func (fc *FacilityController) GetAll(c echo.Context) error {
	facilitiesData := fc.facilityUsecase.GetAll()

	facilities := []response.Facility{}

	for _, facility := range facilitiesData {
		facilities = append(facilities, response.FromDomain(facility))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all facilities", facilities)
}

func (fc *FacilityController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	facility := fc.facilityUsecase.GetByID(id)

	if facility.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "facility not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "facility found", response.FromDomain(facility))
}

func (fc *FacilityController) Create(c echo.Context) error {
	inputTemp := request.Facility{}

	if err := c.Bind(&inputTemp); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	input := request.Facility{
		Description: inputTemp.Description,
	}

	err := input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	facility := fc.facilityUsecase.Create(input.ToDomain())

	return ctrl.NewResponse(c, http.StatusCreated, "success", "facility created", response.FromDomain(facility))
}

func (fc *FacilityController) Update(c echo.Context) error {
	input := request.Facility{}

	if err := c.Bind(&input); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var facilityId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	facility := fc.facilityUsecase.Update(facilityId, input.ToDomain())

	if facility.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "facility not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "facility updated", response.FromDomain(facility))
}

func (fc *FacilityController) Delete(c echo.Context) error {
	var facilityId string = c.Param("id")

	isSuccess := fc.facilityUsecase.Delete(facilityId)

	if !isSuccess {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "facility not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "facility deleted", "")
}
