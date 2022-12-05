package offices

import (
	"backend/businesses/offices"
	"backend/helper"
	"context"
	"fmt"
	"strconv"

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

func (oc *OfficeController) GetAll(c echo.Context) error {
	officesData := oc.officeUsecase.GetAll()

	offices := []response.Office{}

	for _, office := range officesData {
		offices = append(offices, response.FromDomain(office))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all offices", offices)
}

func (oc *OfficeController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	office := oc.officeUsecase.GetByID(id)

	if office.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "office not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "office found", response.FromDomain(office))
}

func (oc *OfficeController) Create(c echo.Context) error {
	ctx := context.Background()
	var imageURLs []string
	var countFiles int
	inputTemp := request.OfficeTemp{}

	if err := c.Bind(&inputTemp); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "bind failed", "")
	}

	openHourTemp, closeHourTemp := helper.ConvertShiftClockToShiftTime(inputTemp.OpenHour, inputTemp.CloseHour)

	// multipart form
	form, err := c.MultipartForm()
	
	if err != nil {
		return err
	}
	
	files := form.File["images"]
	
	for _, file := range files {
		isFileAllowed, isFileAllowedMessage := helper.IsFileAllowed(file)

		if !isFileAllowed {
			return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", isFileAllowedMessage)
		}

		countFiles++
	}

	// limit to only 4 files
	if countFiles != 4 {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "please input four images only")
	}
	
	imageURLs, err = helper.CloudinaryUploadOfficeImgs(ctx, files)

	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusConflict, "failed", "conflict when upload file in cloud image")
	}

	inputTemp.Images = imageURLs

	input := request.Office{
		Title: inputTemp.Title,
		Description: inputTemp.Description,
		OfficeType: inputTemp.OfficeType,
		OfficeLength: inputTemp.OfficeLength,
		PricePerHour: inputTemp.PricePerHour,
		OpenHour: openHourTemp,
		CloseHour: closeHourTemp,
		Lat: inputTemp.Lat,
		Lng: inputTemp.Lng,
		Accommodate: inputTemp.Accommodate,
		WorkingDesk: inputTemp.WorkingDesk,
		MeetingRoom: inputTemp.MeetingRoom,
		PrivateRoom: inputTemp.PrivateRoom,
		City: inputTemp.City,
		District: inputTemp.District,
		Address: inputTemp.Address,
		Images: inputTemp.Images,
	}

	err = input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	office := oc.officeUsecase.Create(input.ToDomain())

	if office.ID == 0 {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "error when creating office")
	}
	
	getOffice := oc.officeUsecase.GetByID(strconv.Itoa(int(office.ID)))

	return ctrl.NewResponse(c, http.StatusCreated, "success", "office created", response.FromDomain(getOffice))
}

func (oc *OfficeController) Update(c echo.Context) error {
	var officeId string = c.Param("id")
	ctx := context.Background()
	inputTemp := request.OfficeTemp{}
	var imageURLs []string

	getOffice := oc.officeUsecase.GetByID(officeId)

	if getOffice.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "office not found", "")
	}

	if err := c.Bind(&inputTemp); err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	openHourTemp, closeHourTemp := helper.ConvertShiftClockToShiftTime(inputTemp.OpenHour, inputTemp.CloseHour)

	// multipart form
	form, err := c.MultipartForm()
	
	if err != nil {
		return err
	}
	
	files := form.File["images"]

	var countFiles int
	
	for _, file := range files {
		isFileAllowed, isFileAllowedMessage := helper.IsFileAllowed(file)

		if !isFileAllowed {
			return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", isFileAllowedMessage)
		}

		countFiles++
	}

	// check if image update available. 
	// if available, limit to only 4 files
	if len(files) != 0 {
		if countFiles != 4 {
			return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "please input four images only")
		}
		
		imageURLs, err = helper.CloudinaryUploadOfficeImgs(ctx, files)
		
		if err != nil {
			return ctrl.NewInfoResponse(c, http.StatusConflict, "failed", "conflict when upload file in cloud image")
		}
		inputTemp.Images = imageURLs
	}
	
	input := request.Office{
		Title: inputTemp.Title,
		Description: inputTemp.Description,
		OfficeType: inputTemp.OfficeType,
		OfficeLength: inputTemp.OfficeLength,
		PricePerHour: inputTemp.PricePerHour,
		OpenHour: openHourTemp,
		CloseHour: closeHourTemp,
		Lat: inputTemp.Lat,
		Lng: inputTemp.Lng,
		Accommodate: inputTemp.Accommodate,
		WorkingDesk: inputTemp.WorkingDesk,
		MeetingRoom: inputTemp.MeetingRoom,
		PrivateRoom: inputTemp.PrivateRoom,
		City: inputTemp.City,
		District: inputTemp.District,
		Address: inputTemp.Address,
		Images: inputTemp.Images,
	}

	err = input.Validate()

	if err != nil {
		return ctrl.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	office := oc.officeUsecase.Update(officeId, input.ToDomain())

	if office.ID == 0 {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "office not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "office updated", response.FromDomain(office))
}

func (oc *OfficeController) Delete(c echo.Context) error {
	var officeId string = c.Param("id")

	isSuccess := oc.officeUsecase.Delete(officeId)

	if !isSuccess {
		return ctrl.NewResponse(c, http.StatusNotFound, "failed", "office not found", "")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "office deleted", "")
}

func (oc *OfficeController) SearchByCity(c echo.Context) error {
	var city string = c.Param("city")

	offices := []response.Office{}

	officesData := oc.officeUsecase.SearchByCity(city)

	for _, office := range officesData {
		offices = append(offices, response.FromDomain(office))
	}

	if len(offices) == 0 {
		return ctrl.NewInfoResponse(c, http.StatusNotFound, "failed", fmt.Sprintf("%s city not found", city))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "grouping by city", offices)
}

func (oc *OfficeController) SearchByRate(c echo.Context) error {
	var rate string = c.Param("rate")
	
	offices := []response.Office{}
	
	officesData := oc.officeUsecase.SearchByRate(rate)

	for _, office := range officesData {
		offices = append(offices, response.FromDomain(office))
	}

	if len(offices) == 0 {
		return ctrl.NewInfoResponse(c, http.StatusNotFound, "failed", fmt.Sprintf("city with rate %s not found", rate))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "grouping by rate", offices)
}

func (oc *OfficeController) SearchByTitle(c echo.Context) error {
	var title string = c.QueryParam("search")

	office := oc.officeUsecase.SearchByTitle(title)

	if office.ID == 0 {
		return ctrl.NewInfoResponse(c, http.StatusNotFound, "failed", "office not found")
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "office found", response.FromDomain(office))
}
