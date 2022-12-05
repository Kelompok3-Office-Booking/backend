package routes

import (
	"backend/controllers/facilities"
	officeimage "backend/controllers/office_images"
	"backend/controllers/offices"
	"backend/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware      echo.MiddlewareFunc
	JWTMiddleware         middleware.JWTConfig
	AuthController        users.AuthController
	OfficeController      offices.OfficeController
	OfficeImageController officeimage.OfficeImageController
	FacilityController    facilities.FacilityController
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

	offices.GET("", cl.OfficeController.GetAll).Name = "get-all-offices"
	offices.GET("/:id", cl.OfficeController.GetByID).Name = "get-office-by-id"
	offices.POST("", cl.OfficeController.Create).Name = "create-office"
	offices.PUT("/:id", cl.OfficeController.Update).Name = "update-office"
	offices.DELETE("/:id", cl.OfficeController.Delete).Name = "delete-office"
	offices.GET("/city/:city", cl.OfficeController.SearchByCity).Name = "group-office-by-city"
	offices.GET("/rate/:rate", cl.OfficeController.SearchByRate).Name = "group-office-by-rate"
	offices.GET("/title", cl.OfficeController.SearchByTitle).Name = "search-office-by-title"
	offices.POST("/office-images", cl.OfficeImageController.Create).Name = "create-office-image-list"

	facilities := e.Group("/api/v1/facilities", middleware.JWTWithConfig(cl.JWTMiddleware))

	facilities.GET("", cl.FacilityController.GetAll).Name = "get-all-facility"
	facilities.GET("/:id", cl.FacilityController.GetByID).Name = "get-facility-by-id"
	facilities.POST("", cl.FacilityController.Create).Name = "create-facility"
	facilities.PUT("/:id", cl.FacilityController.Update).Name = "update-facility"
	facilities.DELETE("/:id", cl.FacilityController.Delete).Name = "delete-facility"

	auth := e.Group("/api/v1", middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.POST("/logout", cl.AuthController.Logout).Name = "user-logout"
}
