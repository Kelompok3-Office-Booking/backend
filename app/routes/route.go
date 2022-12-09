package routes

import (
	"backend/controllers/facilities"
	officefacilities "backend/controllers/office_facilities"
	officeimage "backend/controllers/office_images"
	"backend/controllers/offices"
	transactions "backend/controllers/transactions"
	"backend/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware         echo.MiddlewareFunc
	JWTMiddleware            middleware.JWTConfig
	AuthController           users.AuthController
	OfficeController         offices.OfficeController
	OfficeImageController    officeimage.OfficeImageController
	FacilityController       facilities.FacilityController
	OfficeFacilityController officefacilities.OfficeFacilityController
	TransactionController    transactions.TransactionController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	e.POST("/api/v1/register", cl.AuthController.Register)
	e.POST("/api/v1/login", cl.AuthController.Login)

	users := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	users.GET("", cl.AuthController.GetAll).Name = "get-all-user"
	users.GET("/:id", cl.AuthController.GetByID).Name = "get-user-by-id"
	users.DELETE("/:id", cl.AuthController.Delete).Name = "delete-user-account"
	users.PUT("/profile-photo/:id", cl.AuthController.UpdateProfilePhoto).Name = "update-user-profile-photo"
	users.PUT("/:id", cl.AuthController.UpdateProfileData).Name = "update-profile-data"

	offices := e.Group("/api/v1/offices", middleware.JWTWithConfig(cl.JWTMiddleware))

	offices.GET("/all", cl.OfficeController.GetAll).Name = "get-all-type-of-offices"
	offices.GET("/:id", cl.OfficeController.GetByID).Name = "get-office-by-id"
	offices.POST("/create", cl.OfficeController.Create).Name = "create-office"
	offices.PUT("/update/:office_id", cl.OfficeController.Update).Name = "update-office"
	offices.DELETE("/delete/:office_id", cl.OfficeController.Delete).Name = "delete-office"
	offices.GET("/city/:city", cl.OfficeController.SearchByCity).Name = "group-office-by-city"
	offices.GET("/rate/:rate", cl.OfficeController.SearchByRate).Name = "group-office-by-rate"
	offices.GET("/title", cl.OfficeController.SearchByTitle).Name = "search-office-by-title"
	offices.POST("/images", cl.OfficeImageController.Create).Name = "create-office-image-list"
	offices.GET("/facilities", cl.OfficeFacilityController.GetAll).Name = "get-all-office-facility"
	offices.GET("/facilities/:id", cl.OfficeFacilityController.GetByOfficeID).Name = "get-office-facility-by-id"
	offices.POST("/facilities/create", cl.OfficeFacilityController.Create).Name = "create-office-facility-list"
	offices.GET("/type/office", cl.OfficeController.GetOffices).Name="get-offices"
	offices.GET("/type/coworking-space", cl.OfficeController.GetCoworkingSpace).Name = "get-coworking-spaces"
	offices.GET("/type/meeting-room", cl.OfficeController.GetMeetingRooms).Name = "get-meeting-rooms"
	offices.GET("/recommendation", cl.OfficeController.GetRecommendation).Name = "recommendation-offices"
	offices.GET("/nearest", cl.OfficeController.GetNearest).Name = "get-nearest-building"

	facilities := e.Group("/api/v1/facilities", middleware.JWTWithConfig(cl.JWTMiddleware))

	facilities.GET("", cl.FacilityController.GetAll).Name = "get-all-facility"
	facilities.GET("/:id", cl.FacilityController.GetByID).Name = "get-facility-by-id"
	facilities.POST("", cl.FacilityController.Create).Name = "create-facility"
	facilities.PUT("/:id", cl.FacilityController.Update).Name = "update-facility"
	facilities.DELETE("/:id", cl.FacilityController.Delete).Name = "delete-facility"

	transactions := e.Group("/api/v1/transactions", middleware.JWTWithConfig(cl.JWTMiddleware))

	transactions.GET("", cl.TransactionController.GetAll).Name = "get-all-transaction"
	transactions.POST("", cl.TransactionController.Create).Name = "create-transaction"

	auth := e.Group("/api/v1", middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.POST("/logout", cl.AuthController.Logout).Name = "user-logout"
}
