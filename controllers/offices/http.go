package offices

import (
	"backend/businesses/offices"

	ctrl "backend/controllers"
	"backend/controllers/offices/request"
	"backend/controllers/offices/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type OfficeController struct {
	officeUsecase offices.Usecase
}

func NewOfficeController(officeUC offices.Usecase) *OfficeController {
	return &OfficeController{
		officeUsecase: officeUC,
	}
}

func (oc *OfficeController) Create(c echo.Context) error {
	officeInput := request.Office{}

	if err := c.Bind(&officeInput); err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "invalid request")
	}

	err := officeInput.Validate()
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "validation failed")
	}

	office := oc.officeUsecase.Create(officeInput.ToDomainCreate())

	return ctrl.NewResponse(c, http.StatusCreated, "success", "office created", response.FromDomain(office))
}

func (oc *OfficeController) GetAll(c echo.Context) error {
	offices := []response.Office{}
	officesData := oc.officeUsecase.GetAll()

	for _, office := range officesData {
		offices = append(offices, response.FromDomain(office))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all offices", offices)
}

func (oc *OfficeController) GetByID(c echo.Context) error {
	paramsId := c.Param("id")
	office := oc.officeUsecase.GetByID(paramsId)

	if office.ID == 0 {
		return ctrl.NewInfoResponse(c, http.StatusNotFound, "failed", "office not found")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "office found", response.FromDomain(office))
}

func (oc *OfficeController) Delete(c echo.Context) error {
	paramsId := c.Param("id")
	isSuccess := oc.officeUsecase.Delete(paramsId)

	if !isSuccess {
		return ctrl.NewInfoResponse(c, http.StatusNotFound, "failed", "office not found")
	}

	return ctrl.NewInfoResponse(c, http.StatusOK, "success", "office deleted")
}

func (oc *OfficeController) SearchByCity(c echo.Context) error {
	offices := []response.Office{}
	officesData := oc.officeUsecase.SearchByCity("city")

	for _, office := range officesData {
		offices = append(offices, response.FromDomain(office))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "grouping by city", offices)
}

func (oc *OfficeController) SearchByRate(c echo.Context) error {
	offices := []response.Office{}
	officesData := oc.officeUsecase.SearchByRate("rate")

	for _, office := range officesData {
		offices = append(offices, response.FromDomain(office))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "grouping by rate", offices)
}
