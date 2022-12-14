package main

import (
	_middlewares "backend/app/middlewares"
	_routes "backend/app/routes"
	"os"

	"fmt"

	_driverFactory "backend/drivers"
	_dbDriver "backend/drivers/mysql"

	_userUseCase "backend/businesses/users"
	_userController "backend/controllers/users"

	_officeUseCase "backend/businesses/offices"
	_officeController "backend/controllers/offices"

	_officeImageUseCase "backend/businesses/office_images"
	_officeImageController "backend/controllers/office_images"

	_facilityUseCase "backend/businesses/facilities"
	_facilityController "backend/controllers/facilities"

	_officeFacilityUseCase "backend/businesses/office_facilities"
	_officeFacilityController "backend/controllers/office_facilities"

	_transactionUseCase "backend/businesses/transactions"
	_transactionController "backend/controllers/transactions"

	_reviewUseCase "backend/businesses/review"
	_reviewController "backend/controllers/review"

	"github.com/labstack/echo/v4"
)

const DEFAULT_PORT = "3000"

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT: os.Getenv("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	app := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUseCase)

	officeRepo := _driverFactory.NewOfficeRepository(db)
	officeUseCase := _officeUseCase.NewOfficeUsecase(officeRepo)
	officeCtrl := _officeController.NewOfficeController(officeUseCase)

	officeImageRepo := _driverFactory.NewOfficeImageRepository(db)
	officeImageUseCase := _officeImageUseCase.NewOfficeImageUsecase(officeImageRepo)
	officeImageCtrl := _officeImageController.NewOfficeImageController(officeImageUseCase)

	facilityRepo := _driverFactory.NewFacilityRepository(db)
	facilityUseCase := _facilityUseCase.NewFacilityUsecase(facilityRepo)
	facilityCtrl := _facilityController.NewFacilityController(facilityUseCase)

	officeFacilityRepo := _driverFactory.NewOfficeFacilityRepository(db)
	officeFacilityUseCase := _officeFacilityUseCase.NewOfficeFacilityUsecase(officeFacilityRepo)
	officeFacilityCtrl := _officeFacilityController.NewOfficeFacilityController(officeFacilityUseCase)

	TransactionRepo := _driverFactory.NewTransactionRepository(db)
	TransactionUseCase := _transactionUseCase.NewTransactionUsecase(TransactionRepo)
	TransactionCtrl := _transactionController.NewTransactionController(TransactionUseCase)

	reviewRepo := _driverFactory.NewReviewRepository(db)
	reviewUseCase := _reviewUseCase.NewReviewUsecase(reviewRepo)
	reviewCtrl := _reviewController.NewReviewController(reviewUseCase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:         configLogger.Init(),
		JWTMiddleware:            configJWT.Init(),
		AuthController:           *userCtrl,
		OfficeController:         *officeCtrl,
		OfficeImageController:    *officeImageCtrl,
		FacilityController:       *facilityCtrl,
		OfficeFacilityController: *officeFacilityCtrl,
		TransactionController:    *TransactionCtrl,
		ReviewController: *reviewCtrl,
	}

	routesInit.RouteRegister(app)

	var port string = os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	var appPort string = fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))
}
