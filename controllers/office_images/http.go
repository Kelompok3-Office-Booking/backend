package officeimages

import (
	officeimages "backend/businesses/office_images"
	ctrl "backend/controllers"
	"backend/controllers/office_images/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OfficeImageController struct {
	officeImageUsecase officeimages.Usecase
}

func NewOfficeImageController(uc officeimages.Usecase) *OfficeImageController {
	return &OfficeImageController{
		officeImageUsecase: uc,
	}
}

func (oc *OfficeImageController) GetByOfficeID(c echo.Context) error {
	officeId := c.Param("office_id")

	officeImagesData := oc.officeImageUsecase.GetByOfficeID(officeId)

	officeImages := []response.OfficeImage{}

	for _, v := range officeImagesData {
		officeImages = append(officeImages, response.FromDomain(v))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "get all office images by office_id", officeImages)
}